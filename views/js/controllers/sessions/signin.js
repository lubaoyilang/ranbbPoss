angular.module("app").controller('SessionSigninCtrl', ['$scope', '$modal', '$timeout','$state', function($scope, $modal, $timeout,$state) {
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
            		$scope.login.mode = "success"
            		$scope.login.message = "登陆成功,请稍后...";
            		$timeout(function() {
                  $state.isLogin = true;
            			$state.go('main.dashboard');
            		}, 1000);
            	}

            }, 1000);
        },
        change: function() {
        	$scope.login.mode = "";
            $scope.login.message = "登陆";
        	$scope.login.errors = {}
        }
    }
}]);
