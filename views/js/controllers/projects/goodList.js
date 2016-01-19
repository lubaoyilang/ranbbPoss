angular.module("app").controller('GoodsListCtrl', ['$scope','$modal','$timeout','$http','ngToast', function($scope,$modal,$timeout,$http,ngToast) {

    $scope.modal = {
        open: function(size,class_name) {
            var modalInstance = $modal.open({
                templateUrl: 'modal_add_goods.html',
                controller: 'ModalExampleInstanceCtrl',
                size: size,
                windowClass: class_name
            });

            modalInstance.result.then(function(result) {
            }, function() {
            });
        }
    }

    $scope.project = {
        loading: false,
        page:1,
        size:20,
        scroll: function() {

        }
    }
    $scope.projects = [];
}]);