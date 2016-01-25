/**
 * Created by wangmeng on 16/1/25.
 */
angular.module("app").controller('EditGoodsController', ['$scope','$modal','$timeout','$state','$http','ngToast','$rootScope',function($scope,$modal,$timeout,$state,$http,ngToast,$rootScope) {
    $scope.State = true;
    $scope.goods = $rootScope.goods;
    $scope.Brokerage=  $scope.goods.Brokerage*0.01;
    $scope.SettingPrice = $scope.goods.SettingPrice *0.01;
    $scope.updateGoodsInfo = function(){

        console.log('file is ' );
        console.dir($scope.goods.image);

        var fd = new FormData();
        fd.append('file', $scope.goods.image);
        fd.append("GoodsName",$scope.goods.GoodsName);
        fd.append("Requirelevel",$scope.goods.Requirelevel);
        fd.append("Brokerage",$scope.Brokerage*100);
        fd.append("SettingPrice",$scope.SettingPrice * 100);
        fd.append("Memo",$scope.goods.Memo);
        fd.append("Goodid",$scope.goods.Goodid);
        if($scope.State){
            fd.append("State",1);
        }else{
            fd.append("State",0);
        }

        $http.post('/ranbb/updateGoods', fd, {
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