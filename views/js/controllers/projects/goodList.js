angular.module("app").controller('GoodsListCtrl', ['$rootScope','$scope','$state','$modal','$timeout','$http','ngToast', function($rootScope,$scope,$state,$modal,$timeout,$http,ngToast) {

    $scope.shop = $rootScope.shop;
    $scope.isCollapsed = false;


    $scope.gotoEditGoods = function(goods){
        $rootScope.goods = goods;
        $state.go('main.projects.editGoods');
    }
    $scope.gotoAddGoods = function(){
        $rootScope.shop = $scope.shop;
        $state.go('main.projects.addGoods');
    }
    $scope.gotoCategoryList = function(goods){
        $rootScope.goods = goods;
        $state.go('main.projects.categoryList');
    }

    $scope.project = {
        loading: false,
        size:20,
        scroll: function() {
            if (!$scope.project.loading){
                $scope.project.loading = true;
                $http({method: 'POST',
                    url: '/ranbb/getGoodsList',
                    params:{shopId:$scope.shop.Shopid}
                })
                    .success(function(data, status, headers, config) {
                        console.log(data);
                        if(data.Code > 0&& data.Data != null){
                            $scope.projects=[];
                            Array.prototype.push.apply( $scope.projects, data.Data);
                        }else{
                            ngToast.create({
                                content: data.Data,
                                className: 'warning',
                                dismissButton: true
                            });
                        }
                        $scope.project.loading = false;
                    })
                    .error(function(data, status, headers, config) {
                        ngToast.create({
                            content: "加载失败",
                            className: 'warning',
                            dismissButton: true
                        });
                        $scope.project.loading = false;
                    });
            }
        }
    }
    $scope.projects = [];
}]);