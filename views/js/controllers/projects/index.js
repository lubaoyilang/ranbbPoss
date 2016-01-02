angular.module("app").controller('ProjectIndexCtrl', ['$scope','$modal','$timeout', function($scope,$modal,$timeout) {

    $scope.modal = {
        open: function(size,class_name) {
            var modalInstance = $modal.open({
                templateUrl: 'modal_example.html',
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
                        name: "新加载的店1",
                        icon: "chronometer",
                        color: "#3498DB",
                        progress: {
                            percentage: 80,
                            status: "warning"
                        }
                    }, {
                        name: "新加载的店2",
                        icon: "screen79",
                        color: "#1ABC9C",
                        progress: {
                            percentage: 40,
                            status: "danger"
                        }
                    }, {
                        name: "新加载的店3",
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
            name: "361童鞋天猫旗舰店1",
            icon: "chronometer",
            color: "#3498DB",
            progress: {
                percentage: 80,
                status: "warning"
            }
        },
        {
            name: "361童鞋天猫旗舰店2",
            icon: "screen79",
            color: "#1ABC9C",
            progress: {
                percentage: 40,
                status: "danger"
            }
        },
        {
            name: "361童鞋天猫旗舰店2",
            icon: "objective",
            color: "#F04903",
            progress: {
                percentage: 96,
                status: "success"
            }
        },{
            name: "361童鞋天猫旗舰店3",
            icon: "chronometer",
            color: "#3498DB",
            progress: {
                percentage: 80,
                status: "warning"
            }
        },
        {
            name: "361童鞋天猫旗舰店4",
            icon: "screen79",
            color: "#1ABC9C",
            progress: {
                percentage: 40,
                status: "danger"
            }
        },
        {
            name: "361童鞋天猫旗舰店5",
            icon: "objective",
            color: "#F04903",
            progress: {
                percentage: 96,
                status: "success"
            }
        },{
            name: "361童鞋天猫旗舰店6",
            icon: "chronometer",
            color: "#3498DB",
            progress: {
                percentage: 80,
                status: "warning"
            }
        }
    ];
}]);
