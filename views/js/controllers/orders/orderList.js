angular.module("app").controller('OrderListController', ['$scope','$modal','$timeout','$state','$http','ngToast','$rootScope',function($scope,$modal,$timeout,$state,$http,ngToast,$rootScope) {
    $scope.goods = $rootScope.goods;

    $scope.searchText = "";

    $scope.todo = {
        loading: false,
        page:1,
        size:0,
        scroll: function() {
            if (!$scope.todo.loading) {
                $scope.todo.loading = true;
                $http({method: 'POST',
                    url: '/ranbb/getOrders',
                    params:{goodsId:$scope.goods.Goodid,page:$scope.todo.page}
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