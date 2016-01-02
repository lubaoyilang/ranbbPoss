angular.module("app").controller('ElementInfiniteScrollTableCtrl', ['$scope', '$modal', '$timeout', function($scope, $modal, $timeout) {

    $scope.todo = {
        loading: false,
        scroll: function() {
            if (!$scope.todo.loading) {
                $scope.todo.loading = true;
                $timeout(function() {
                    $scope.todo.list.push({
                        title: "Lorem ipsum dolor sit amet, consectetuer adipiscing elit.",
                        date: new Date(),
                        status: "danger",
                    }, {
                        title: "Aliquam tincidunt mauris eu risus.",
                        date: new Date(),
                        status: "danger"
                    }, {
                        title: "Vestibulum auctor dapibus neque.",
                        date: new Date(),
                        status: "danger"
                    }, {
                        title: "Nunc dignissim risus id metus.",
                        date: new Date(),
                        status: "danger"
                    }, {
                        title: "Cras ornare tristique elit.",
                        date: new Date(),
                        status: "danger"
                    }, {
                        title: "Vivamus vestibulum nulla nec ante.",
                        date: new Date(),
                        status: "danger"
                    }, {
                        title: "Praesent placerat risus quis eros.",
                        date: new Date(),
                        status: "warning"
                    }, {
                        title: "Fusce pellentesque suscipit nibh.",
                        date: new Date(),
                        status: "warning"
                    });
                    $scope.todo.loading = false;
                }, 1000);
            }
        },
        list: [{
            title: "Lorem ipsum dolor sit amet, consectetuer adipiscing elit.",
            date: new Date(),
            status: "danger",
        }, {
            title: "Aliquam tincidunt mauris eu risus.",
            date: new Date(),
            status: "danger"
        }, {
            title: "Vestibulum auctor dapibus neque.",
            date: new Date(),
            status: "danger"
        }, {
            title: "Nunc dignissim risus id metus.",
            date: new Date(),
            status: "danger"
        }, {
            title: "Cras ornare tristique elit.",
            date: new Date(),
            status: "danger"
        }, {
            title: "Vivamus vestibulum nulla nec ante.",
            date: new Date(),
            status: "danger"
        }, {
            title: "Praesent placerat risus quis eros.",
            date: new Date(),
            status: "warning"
        }, {
            title: "Fusce pellentesque suscipit nibh.",
            date: new Date(),
            status: "warning"
        }, {
            title: "Lorem ipsum dolor sit amet, consectetuer adipiscing elit.",
            date: new Date(),
            status: "danger",
        }, {
            title: "Aliquam tincidunt mauris eu risus.",
            date: new Date(),
            status: "danger"
        }, {
            title: "Vestibulum auctor dapibus neque.",
            date: new Date(),
            status: "danger"
        }, {
            title: "Nunc dignissim risus id metus.",
            date: new Date(),
            status: "danger"
        }, {
            title: "Cras ornare tristique elit.",
            date: new Date(),
            status: "danger"
        }, {
            title: "Vivamus vestibulum nulla nec ante.",
            date: new Date(),
            status: "danger"
        }, {
            title: "Praesent placerat risus quis eros.",
            date: new Date(),
            status: "warning"
        }, {
            title: "Fusce pellentesque suscipit nibh.",
            date: new Date(),
            status: "warning"
        }, {
            title: "Lorem ipsum dolor sit amet, consectetuer adipiscing elit.",
            date: new Date(),
            status: "danger",
        }, {
            title: "Aliquam tincidunt mauris eu risus.",
            date: new Date(),
            status: "danger"
        }, {
            title: "Vestibulum auctor dapibus neque.",
            date: new Date(),
            status: "danger"
        }, {
            title: "Nunc dignissim risus id metus.",
            date: new Date(),
            status: "danger"
        }, {
            title: "Cras ornare tristique elit.",
            date: new Date(),
            status: "danger"
        }, {
            title: "Vivamus vestibulum nulla nec ante.",
            date: new Date(),
            status: "danger"
        }, {
            title: "Praesent placerat risus quis eros.",
            date: new Date(),
            status: "warning"
        }, {
            title: "Fusce pellentesque suscipit nibh.",
            date: new Date(),
            status: "warning"
        }]
    };
}]);
