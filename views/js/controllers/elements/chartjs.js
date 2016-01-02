angular.module("app").controller('ElementChartJSCtrl', ['$scope', 'DTOptionsBuilder', 'DTColumnBuilder', function($scope, DTOptionsBuilder, DTColumnBuilder) {
    $scope.labels = ["January", "February", "March", "April", "May", "June", "July"];
    $scope.series = ['Series A', 'Series B', 'Series C'];
    $scope.colours = [{
        fillColor: "rgba(26,188,156,1)",
        pointColor: "rgba(26,188,156,1)",
        pointHighlightFill: "#fff",
        pointHighlightStroke: "rgba(26,188,156,0.8)",
        pointStrokeColor: "#fff",
        strokeColor: "rgba(26,188,156,0)"
    }, {
        fillColor: "rgba(96,100,109,1)",
        pointColor: "rgba(96,100,109,1)",
        pointHighlightFill: "#fff",
        pointHighlightStroke: "rgba(96,100,109,0.8)",
        pointStrokeColor: "#fff",
        strokeColor: "rgba(96,100,109,0)"
    }, {
        fillColor: "rgba(52,152,219,1)",
        pointColor: "rgba(52,152,219,1)",
        pointHighlightFill: "#fff",
        pointHighlightStroke: "rgba(52,152,219,0.8)",
        pointStrokeColor: "#fff",
        strokeColor: "rgba(52,152,219,0)"
    }];

    $scope.data = [
        [65, 59, 80, 81, 56, 55, 40],
        [28, 48, 40, 19, 86, 27, 90],
        [103, 20, 20, 19, 75, 22, 10]
    ];
    $scope.pie_labels = ["Download Sales", "In-Store Sales", "Mail-Order Sales"];
    $scope.pie_data = [300, 500, 100];

    $scope.onClick = function(points, evt) {
        console.log(points, evt);
    };
}]);
