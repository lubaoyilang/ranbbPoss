angular.module("app").controller('DashboardCtrl', ['$scope','$state',function($scope,$state) {

    $scope.todo = {
        list: [{
            title: "马培龙",
            date: new Date(),
            status: "danger",
        },{
            title: "马培龙",
            date: new Date(),
            status: "danger"
        },{
            title: "马培龙",
            date: new Date(),
            status: "danger"
        },{
            title: "马培龙",
            date: new Date(),
            status: "danger"
        },{
            title: "马培龙",
            date: new Date(),
            status: "danger"
        },{
            title: "马培龙",
            date: new Date(),
            status: "danger"
        },{
            title: "马培龙",
            date: new Date(),
            status: "warning"
        },{
            title: "马培龙",
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
            name: "童鞋旗舰店1",
            icon: "chronometer",
            color: "#3498DB",
            progress: {
                percentage: 80,
                status: "warning"
            }
        },
        {
            name: "童鞋旗舰店2",
            icon: "screen79",
            color: "#1ABC9C",
            progress: {
                percentage: 40,
                status: "danger"
            }
        },
        {
            name: "童鞋旗舰店3",
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
