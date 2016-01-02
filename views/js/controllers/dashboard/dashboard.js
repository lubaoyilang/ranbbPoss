angular.module("app").controller('DashboardCtrl', ['$scope','$state',function($scope,$state) {
    if ($state.isLogin == false) {
      $state.go("sessions.signin")
    }
    
    $scope.todo = {
        list: [{
            title: "Lorem ipsum dolor sit amet, consectetuer adipiscing elit.",
            date: new Date(),
            status: "danger",
        },{
            title: "Aliquam tincidunt mauris eu risus.",
            date: new Date(),
            status: "danger"
        },{
            title: "Vestibulum auctor dapibus neque.",
            date: new Date(),
            status: "danger"
        },{
            title: "Nunc dignissim risus id metus.",
            date: new Date(),
            status: "danger"
        },{
            title: "Cras ornare tristique elit.",
            date: new Date(),
            status: "danger"
        },{
            title: "Vivamus vestibulum nulla nec ante.",
            date: new Date(),
            status: "danger"
        },{
            title: "Praesent placerat risus quis eros.",
            date: new Date(),
            status: "warning"
        },{
            title: "Fusce pellentesque suscipit nibh.",
            date: new Date(),
            status: "warning"
        }],
        add: function(task) {
            console.log(task);
        }
    };

    $scope.project = {};
    $scope.projects = [
        {
            name: "Message Backend",
            icon: "chronometer",
            color: "#3498DB",
            progress: {
                percentage: 80,
                status: "warning"
            }
        },
        {
            name: "Location-Base iOS App",
            icon: "screen79",
            color: "#1ABC9C",
            progress: {
                percentage: 40,
                status: "danger"
            }
        },
        {
            name: "Company Goal",
            icon: "objective",
            color: "#F04903",
            progress: {
                percentage: 96,
                status: "success"
            }
        }
    ];


    $scope.pie_labels = ["Design", "Development", "Testing"];
    $scope.pie_data = [300, 500, 100];
}]);
