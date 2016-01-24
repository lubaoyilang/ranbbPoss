/**
 * Created by wangmeng on 16/1/23.
 */
angular.module("app").controller('EditShopController', ['$scope','$modal','$timeout','$state','$http','ngToast','$rootScope',function($scope,$modal,$timeout,$state,$http,ngToast,$rootScope) {
    $scope.shop = $rootScope.shop;
    $scope.updateShopInfo = function(){

        console.log('file is ' );
        console.dir($scope.shop.image);

        var fd = new FormData();
        fd.append('file', $scope.shop.image);
        fd.append("Shopid",$scope.shop.Shopid);
        fd.append("Shopname",$scope.shop.Shopname);
        fd.append("Mobile",$scope.shop.Mobile);
        fd.append("Email",$scope.shop.Email);
        fd.append("Shoptaobaoid",$scope.shop.Shoptaobaoid);
        fd.append("Memo",$scope.shop.Memo);
        $http.post('/ranbb/updateShop', fd, {
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