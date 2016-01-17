/**
 * Created by wangmeng on 16/1/17.
 */
angular.module('app').controller('adminEditController', function ($scope,$http,$modalInstance) {

    $scope.admin = $modalInstance.user;
    $scope.admin.Password = "";

    $scope.ok = function () {
        if($scope.state){
            $scope.admin.State = 1;
        }else{
            $scope.admin.State = 0;
        }
        $modalInstance.close($scope.admin);
    };

    $scope.cancel = function () {
        $modalInstance.dismiss('cancel');
    };
});