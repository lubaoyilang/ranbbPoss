/**
 * Created by demon on 16/1/28.
 */
angular.module("app").controller('WalletController', ['$scope','$modal','$timeout','$state','$http','ngToast','$rootScope',function($scope,$modal,$timeout,$state,$http,ngToast,$rootScope) {

    $scope.enchash = function(enchashment){
        $http({method: 'POST',
            url: '/ranbb/enchash',
            params:{enchashId:enchashment.Id}
        })
            .success(function(data, status, headers, config) {
                console.log(data);
                if(data.Code > 0&& data.Data != null){
                    ngToast.create({
                        content: data.Data,
                        className: 'warning',
                        dismissButton: true
                    });
                    enchashment.State = 1;
                }else{
                    ngToast.create({
                        content: data.Data,
                        className: 'warning',
                        dismissButton: true
                    });
                }
                $scope.todo.loading = false;
            })
            .error(function(data, status, headers, config) {
                ngToast.create({
                    content: "加载失败",
                    className: 'warning',
                    dismissButton: true
                });
            });
    }
    $scope.cannotchash = function(enchashment){
        $http({method: 'POST',
            url: '/ranbb/canNotEnchash',
            params:{enchashId:enchashment.Id}
        })
            .success(function(data, status, headers, config) {
                console.log(data);
                if(data.Code > 0&& data.Data != null){
                    ngToast.create({
                        content: data.Data,
                        className: 'warning',
                        dismissButton: true
                    });
                    enchashment.State = 1;
                }else{
                    ngToast.create({
                        content: data.Data,
                        className: 'warning',
                        dismissButton: true
                    });
                }
                $scope.todo.loading = false;
            })
            .error(function(data, status, headers, config) {
                ngToast.create({
                    content: "加载失败",
                    className: 'warning',
                    dismissButton: true
                });
            });
    }
    $scope.todo = {
        loading: false,
        page:1,
        size:0,
        scroll: function() {
            if (!$scope.todo.loading) {
                $scope.todo.loading = true;
                $http({method: 'POST',
                    url: '/ranbb/getAllEnchashments',
                    params:{page:$scope.todo.page}
                })
                    .success(function(data, status, headers, config) {
                        console.log(data);
                        if(data.Code > 0&& data.Data != null){
                            $scope.todo.page++;
                            Array.prototype.push.apply( $scope.todo.list, data.Data);
                            $scope.todo.size = $scope.todo.list.length;
                        }else{
                            ngToast.create({
                                content: data.Data,
                                className: 'warning',
                                dismissButton: true
                            });
                        }
                        $scope.todo.loading = false;
                    })
                    .error(function(data, status, headers, config) {
                        ngToast.create({
                            content: "加载失败",
                            className: 'warning',
                            dismissButton: true
                        });
                        $scope.todo.loading = false;
                    });
            }
        },
        list: []
    };

}]);