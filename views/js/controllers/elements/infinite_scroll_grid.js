angular.module("app").controller('ElementInfiniteScrollGridCtrl', ['$scope', '$modal','$timeout', function($scope, $modal,$timeout) {
    $scope.project = {
        loading: false,
        scroll: function() {
            if (!$scope.project.loading) {
                $scope.project.loading = true;
                $timeout(function() {
		            $scope.projects.push({
		                name: "Message Backend",
		                icon: "chronometer",
		                color: "#3498DB",
		                progress: {
		                    percentage: 80,
		                    status: "warning"
		                }
		            }, {
		                name: "Location-Base iOS App",
		                icon: "screen79",
		                color: "#1ABC9C",
		                progress: {
		                    percentage: 40,
		                    status: "danger"
		                }
		            }, {
		                name: "Company Goal",
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
    $scope.projects = [{
        name: "Message Backend",
        icon: "chronometer",
        color: "#3498DB",
        progress: {
            percentage: 80,
            status: "warning"
        }
    }, {
        name: "Location-Base iOS App",
        icon: "screen79",
        color: "#1ABC9C",
        progress: {
            percentage: 40,
            status: "danger"
        }
    }, {
        name: "Company Goal",
        icon: "objective",
        color: "#F04903",
        progress: {
            percentage: 96,
            status: "success"
        }
    }, {
        name: "Message Backend",
        icon: "chronometer",
        color: "#3498DB",
        progress: {
            percentage: 80,
            status: "warning"
        }
    }, {
        name: "Location-Base iOS App",
        icon: "screen79",
        color: "#1ABC9C",
        progress: {
            percentage: 40,
            status: "danger"
        }
    }, {
        name: "Company Goal",
        icon: "objective",
        color: "#F04903",
        progress: {
            percentage: 96,
            status: "success"
        }
    }, {
        name: "Message Backend",
        icon: "chronometer",
        color: "#3498DB",
        progress: {
            percentage: 80,
            status: "warning"
        }
    }, {
        name: "Location-Base iOS App",
        icon: "screen79",
        color: "#1ABC9C",
        progress: {
            percentage: 40,
            status: "danger"
        }
    }, {
        name: "Company Goal",
        icon: "objective",
        color: "#F04903",
        progress: {
            percentage: 96,
            status: "success"
        }
    }, {
        name: "Message Backend",
        icon: "chronometer",
        color: "#3498DB",
        progress: {
            percentage: 80,
            status: "warning"
        }
    }, {
        name: "Location-Base iOS App",
        icon: "screen79",
        color: "#1ABC9C",
        progress: {
            percentage: 40,
            status: "danger"
        }
    }, {
        name: "Company Goal",
        icon: "objective",
        color: "#F04903",
        progress: {
            percentage: 96,
            status: "success"
        }
    }, {
        name: "Message Backend",
        icon: "chronometer",
        color: "#3498DB",
        progress: {
            percentage: 80,
            status: "warning"
        }
    }, {
        name: "Location-Base iOS App",
        icon: "screen79",
        color: "#1ABC9C",
        progress: {
            percentage: 40,
            status: "danger"
        }
    }, {
        name: "Company Goal",
        icon: "objective",
        color: "#F04903",
        progress: {
            percentage: 96,
            status: "success"
        }
    }, {
        name: "Message Backend",
        icon: "chronometer",
        color: "#3498DB",
        progress: {
            percentage: 80,
            status: "warning"
        }
    }, {
        name: "Location-Base iOS App",
        icon: "screen79",
        color: "#1ABC9C",
        progress: {
            percentage: 40,
            status: "danger"
        }
    }, {
        name: "Company Goal",
        icon: "objective",
        color: "#F04903",
        progress: {
            percentage: 96,
            status: "success"
        }
    }, {
        name: "Message Backend",
        icon: "chronometer",
        color: "#3498DB",
        progress: {
            percentage: 80,
            status: "warning"
        }
    }, {
        name: "Location-Base iOS App",
        icon: "screen79",
        color: "#1ABC9C",
        progress: {
            percentage: 40,
            status: "danger"
        }
    }, {
        name: "Company Goal",
        icon: "objective",
        color: "#F04903",
        progress: {
            percentage: 96,
            status: "success"
        }
    }, {
        name: "Message Backend",
        icon: "chronometer",
        color: "#3498DB",
        progress: {
            percentage: 80,
            status: "warning"
        }
    }, {
        name: "Location-Base iOS App",
        icon: "screen79",
        color: "#1ABC9C",
        progress: {
            percentage: 40,
            status: "danger"
        }
    }, {
        name: "Company Goal",
        icon: "objective",
        color: "#F04903",
        progress: {
            percentage: 96,
            status: "success"
        }
    }];
}]);
