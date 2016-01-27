/**
 * Created by wangmeng on 16/1/26.
 */
angular.module("app").controller('AddCategoryController', ['$scope','$modal','$timeout','$state','$http','ngToast','$rootScope',function($scope,$modal,$timeout,$state,$http,ngToast,$rootScope) {
    $scope.goods = $rootScope.goods;
    $scope.enableTime = 0;
    $scope.Price  = 0;
    var nowDate = new Date();
    nowDate.setMilliseconds(0);
    nowDate.setSeconds(0);
    $scope.dt = new Date((nowDate.getTime()/1000).toFixed(0)*1000);
    $scope.category = {
        Goodid:$scope.goods.Goodid,
        Shopid:$scope.goods.Shopid};
    $scope.addCategory = function(){

        var fd = new FormData();
        fd.append('Goodid', $scope.goods.Goodid);
        fd.append("Name",$scope.category.Name);
        //fd.append("EnableTime",$scope.category.EnableTime);
        fd.append("TotalNum",$scope.category.TotalNum);
        fd.append("Price",$scope.Price * 100);
        fd.append("Memo",$scope.category.Memo);
        fd.append("GoodsId",$scope.category.Goodid);
        fd.append("ShopId",$scope.category.Shopid);
        fd.append("EnableTime",$scope.dt.getTime()/1000);

        $http.post('/ranbb/addCategory',fd, {
            transformRequest: angular.identity,
            headers: {'Content-Type': undefined}
        }).success(function(data){
                if(data != undefined){
                    if(data.Code < 0){
                        ngToast.create({
                            content: data.Data,
                            className: 'warning',
                            dismissButton: true
                        });
                    }else{
                        window.history.back(-1);
                    }
                }else{
                    ngToast.create({
                        content: "错误的信息",
                        className: 'warning',
                        dismissButton: true
                    });
                }
            })
            .error(function(data){
                ngToast.create({
                    content: data.Data,
                    className: 'warning',
                    dismissButton: true
                });
            });
    }
    $scope.backToShopList = function(){
        window.history.back(-1);
    }
}]);