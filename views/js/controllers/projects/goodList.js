angular.module("app").controller('GoodsListCtrl', ['$scope','$modal','$timeout', function($scope,$modal,$timeout) {

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
        scroll: function() {
            if (!$scope.project.loading) {
                $scope.project.loading = true;
                $timeout(function() {
                    $scope.projects.push({
                        name: "童鞋10块钱一件,包邮!",
                        icon: "chronometer",
                        color: "#3498DB",
                        progress: {
                            percentage: 80,
                            status: "warning"
                        }
                    }, {
                        name: "童鞋10块钱一件,包邮!",
                        icon: "screen79",
                        color: "#1ABC9C",
                        progress: {
                            percentage: 40,
                            status: "danger"
                        }
                    }, {
                        name: "童鞋10块钱一件,包邮!",
                        icon: "objective",
                        color: "#F04903",
                        progress: {
                            percentage: 96,
                            status: "success"
                        }
                    });
                    $scope.project.loading = false;
                }, 1000);
            }
        }
    }
    $scope.projects = [
        {
            name: "童鞋10块钱一件,包邮!",
            icon: "chronometer",
            color: "#3498DB",
            progress: {
                percentage: 80,
                status: "warning"
            }
        },
        {
            name: "童鞋10块钱一件,包邮!",
            icon: "screen79",
            color: "#1ABC9C",
            progress: {
                percentage: 40,
                status: "danger"
            }
        },
        {
            name: "童鞋10块钱一件,包邮!",
            icon: "objective",
            color: "#F04903",
            progress: {
                percentage: 96,
                status: "success"
            }
        },{
            name: "童鞋10块钱一件,包邮!",
            icon: "chronometer",
            color: "#3498DB",
            progress: {
                percentage: 80,
                status: "warning"
            }
        },
        {
            name: "童鞋10块钱一件,包邮!",
            icon: "screen79",
            color: "#1ABC9C",
            progress: {
                percentage: 40,
                status: "danger"
            }
        },
        {
            name: "童鞋10块钱一件,包邮!",
            icon: "objective",
            color: "#F04903",
            progress: {
                percentage: 96,
                status: "success"
            }
        },{
            name: "童鞋10块钱一件,包邮!",
            icon: "chronometer",
            color: "#3498DB",
            progress: {
                percentage: 80,
                status: "warning"
            }
        }
    ];
}]);