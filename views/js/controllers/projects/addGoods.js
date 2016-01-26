/**
 * Created by wangmeng on 16/1/26.
 */
angular.module("app").controller('AddGoodsController', ['$scope','$modal','$timeout','$state','$http','ngToast','$rootScope',function($scope,$modal,$timeout,$state,$http,ngToast,$rootScope) {
    $scope.shop = $rootScope.shop;
    $scope.State = true;
    $scope.goods = {};
    $scope.Brokerage= 0;// $scope.goods.Brokerage*0.01;
    $scope.SettingPrice = 0;//$scope.goods.SettingPrice *0.01;
    $scope.addGoods = function(){

        console.log('file is ' );
        console.dir($scope.goods.image);

        var fd = new FormData();
        fd.append('file', $scope.goods.image);
        fd.append("GoodsName",$scope.goods.GoodsName);
        fd.append("Requirelevel",$scope.goods.Requirelevel);
        fd.append("Brokerage",$scope.Brokerage*100);
        fd.append("SettingPrice",$scope.SettingPrice * 100);
        fd.append("Memo",$scope.goods.Memo);
        fd.append("Shopid",$scope.shop.Shopid);
        if($scope.State){
            fd.append("State",1);
        }else{
            fd.append("State",0);
        }

        $http.post('/ranbb/addGoods', fd, {
            transformRequest: angular.identity,
            headers: {'Content-Type': undefined}
        })
            .success(function(data){
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
                        content: data.Data,
                        className: 'warning',
                        dismissButton: true
                    });
                }
            })
            .error(function(){
            });
    };
    $scope.backToShopList = function(){
        window.history.back(-1);
    }

}]);
angular.module("app").directive('fileModel', ['$parse', function ($parse) {
    return {
        restrict: 'A',
        link: function(scope, element, attrs) {
            var model = $parse(attrs.fileModel);
            var modelSetter = model.assign;

            element.bind('change', function(){
                scope.$apply(function(){
                    modelSetter(scope, element[0].files[0]);
                });
            });
        }
    };
}]);