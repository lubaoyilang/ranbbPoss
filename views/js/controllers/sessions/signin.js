angular.module("app").controller('SessionSigninCtrl', ['$scope', '$modal', '$timeout','$state','$http', function($scope, $modal, $timeout,$state,$http) {
  $state.isLogin = false;
    $scope.login = {
        mode: "",
        message: "登陆",
        errors: {
        	username: false,
        	password: false
        },
        user: {
            username: "",
            password: ""
        },
        login: function(user) {
        	$scope.login.errors = {}
            $scope.login.mode = "loading";
            $scope.login.message = "正在登陆,请稍后...";
            $timeout(function() {
            	if(user.username == "" && user.password == "") {
            		$scope.login.message = "请输入用户名密码";
            		$scope.login.errors.username = true;
            		$scope.login.errors.password = true;
            		$scope.login.mode = "errors"
            	} else if (user.username == "") {
            		$scope.login.message = "请输入用户名";
            		$scope.login.errors.username = true;
            		$scope.login.mode = "errors"
            	} else if (user.password == "") {
            		$scope.login.message = "请输入密码";
            		$scope.login.errors.password = true;
            		$scope.login.mode = "errors"
            	} else {
            		// $scope.login.mode = "success"
            		// $scope.login.message = "登陆成功,请稍后...";
            		// $timeout(function() {
                //   $state.isLogin = true;
            		// 	$state.go('main.dashboard');
            		// }, 1000);

                $http({method: 'POST',
                      url: '/ranbb/login',
                      params:{'username':user.username,'password':user.password}})
                .success(function(data, status, headers, config) {
                      console.log(data);
                      if (data.username == user.username) {
                          $state.isLogin = true; 
                          $state.go('main.dashboard');
                      }else{
                        $scope.login.mode = "errors";
                    		$scope.login.message = data;
                      }
                })
                .error(function(data, status, headers, config) {
                        scope.login.mode = "errors";
                        $scope.login.message = "网络错误";
                });
            	}

            }, 10);
        },
        change: function() {
        	$scope.login.mode = "";
            $scope.login.message = "登陆";
        	$scope.login.errors = {}
        }
    }
}]);
