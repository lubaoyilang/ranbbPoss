angular.module('app').controller('ModalExampleInstanceCtrl', function ($scope, $modalInstance) {

  $scope.shop = $modalInstance.shop;

  $scope.ok = function () {
    $modalInstance.close('ok');
  };

  $scope.cancel = function () {
    $modalInstance.dismiss('cancel');
  };
});