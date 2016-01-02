angular.module("app").controller('ElementComponentsCtrl', ['$scope','$modal','ngToast', function($scope,$modal,ngToast) {
    
    $scope.tabs = {
        classSelected: "",
        classOptions: [
            {
                name: "Primary",
                class_name: "navs-primary"
            },
            {
                name: "Success",
                class_name: "navs-success"
            },
        ]
    }

    $scope.modal = {
        open: function(size,class_name) {
            var modalInstance = $modal.open({
                templateUrl: 'modal_example.html',
                controller: 'ModalExampleInstanceCtrl',
                size: size,
                windowClass: class_name
            });

            modalInstance.result.then(function(result) {
                alert(result);
            }, function() {
                alert("asdf");
            });
        }
    }

    $scope.toast = {
        create: function(message,className) {
            ngToast.create({
                content: message,
                className: className,
                dismissButton: true
            });
        },
        clear: function() {
            ngToast.dismiss();
        }
    }

    $scope.dynamicPopover = 'Hello, World!';
    $scope.dynamicPopoverTitle = 'Title';
}]);
