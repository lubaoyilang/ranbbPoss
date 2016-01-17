/**
 * Created by wangmeng on 16/1/17.
 */
angular.module("app").controller('CreateAdminController', ['$scope', '$modal', '$timeout','$http','$state','ngToast', function($scope, $modal, $timeout,$http,$state,ngToast) {
    $scope.user = {
        Mobile:"",
        Name:"",
        Password:"",
        Privilege:0,
        error:false,
        errorMsg:"",
        addUser:function(){
            $http({method: 'POST',
                url: '/ranbb/addAdmin',
                params:$scope.user
            })
                .success(function(data, status, headers, config) {
                    console.log(data);
                    if(data==null){
                        $scope.user.error = true;
                        $scope.user.errorMsg = "手机号或用户名已经被使用"
                        ngToast.create({
                            content: "手机号或用户名已经被使用",
                            className: 'warning',
                            dismissButton: true
                        });
                        return
                    }
                    window.history.back(-1);
                })
                .error(function(data, status, headers, config) {
                    $scope.user.error = true;
                    $scope.user.errorMsg = "添加失败"
                    ngToast.create({
                        content: "添加失败",
                        className: 'error',
                        dismissButton: true
                    });
                });
        },
        cancel:function(){
            window.history.back(-1);
        }
    }
}]);