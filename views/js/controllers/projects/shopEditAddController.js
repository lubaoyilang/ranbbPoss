/**
 * Created by wangmeng on 16/1/18.
 */
angular.module('app').controller('ShopEditAndAddController', function ($scope,$http,$modalInstance) {
    $scope.shop = $modalInstance.shop;
    $scope.selectedImage = function(inputFile){
        $scope.shop.image = inputFile.value;
    }

    $scope.ok = function () {

        var file=$scope.shop.image;
        $http({
            method: 'POST',
            url: '/ranbb/updateShop',
            headers: { 'Content-Type': 'multipart/form-data'},
            data: $.param($scope.formData)
        })
            .success(function(data, status) {
                alert("Success ... " + status);
            })
            .error(function(data, status) {
                alert("Error ... " + status);
            });


        $modalInstance.close('ok');
    };

    $scope.cancel = function () {
        $modalInstance.dismiss('cancel');
    };
});