/**
 * Created by wangmeng on 16/1/26.
 */
angular.module("app").controller('EditCategoryController', ['$scope','$modal','$timeout','$state','$http','ngToast','$rootScope',function($scope,$modal,$timeout,$state,$http,ngToast,$rootScope) {
    $scope.goods = $rootScope.goods;
    $scope.category = $rootScope.category;
    $scope.Price  = $scope.category.Price/100;
    $scope.dt = new Date($scope.category.EnableTime*1000);
    $scope.updateCategory = function(){
        var fd = new FormData();
        fd.append("Name",$scope.category.Name);
        //fd.append("EnableTime",$scope.category.EnableTime);
        fd.append("TotalNum",$scope.category.TotalNum);
        fd.append("Price",$scope.Price * 100);
        fd.append("Memo",$scope.category.Memo);
        fd.append("EnableTime",$scope.dt.getTime()/1000);
        fd.append("CategroyId",$scope.category.CategroyId);

        $http.post('/ranbb/editCategory',fd, {
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