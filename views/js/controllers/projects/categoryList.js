angular.module("app").controller('CategoryListController', ['$scope','$modal','$timeout','$state','$http','ngToast','$rootScope',function($scope,$modal,$timeout,$state,$http,ngToast,$rootScope) {
    $scope.goods = $rootScope.goods;

    $scope.searchText = "";

    $scope.getRate = function(a,b){
        var c = a/b*100;
        return c.toFixed(2);
    }

    $scope.addCategory = function(){
        $rootScope.goods = $scope.goods;
        $state.go('main.projects.addCategory');
    }

    $scope.gotoOrderList = function(){
        $rootScope.goods = $scope.goods;
        $state.go('main.projects.orderList');
    }

    $scope.editCategory = function(category){
        $rootScope.category = category;
        $state.go('main.projects.editCategory');
    }

    $scope.deleteCategory = function(category){
        $http({method: 'POST',
            url: '/ranbb/deleteCategory',
            params:{categoryId:category.CategroyId}
        })
            .success(function(data, status, headers, config) {
                console.log(data);
                if(data.Code > 0&& data.Data != null){
                    var index = $scope.todo.list.indexOf(category);
                    $scope.todo.list.splice(index,1);
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
                    content: "操作失败",
                    className: 'warning',
                    dismissButton: true
                });
                $scope.todo.loading = false;
            });
    }


    $scope.todo = {
        loading: false,
        size:0,
        scroll: function() {
            if (!$scope.todo.loading) {
                $scope.todo.loading = true;
                $http({method: 'POST',
                    url: '/ranbb/getCategorys',
                    params:{goodsId:$scope.goods.Goodid}
                })
                    .success(function(data, status, headers, config) {
                        console.log(data);
                        if(data.Code > 0&& data.Data != null){
                            $scope.todo.list=[];
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