angular.module("app").controller('DashboardCtrl', ['$scope','$state','$http',function($scope,$state,$http) {



    //获取最新店铺
    $http({method: 'POST',
        url: '/ranbb/newShops'
        })
        .success(function(data, status, headers, config) {
            console.log(data);
            $scope.projects = data
        })
        .error(function(data, status, headers, config) {

        });
    //获取最新订单
    $http({method: 'POST',
        url: '/ranbb/neworders'
    }).success(function(data, status, headers, config) {
            console.log(data);
            $scope.todo = {list:data}
     }).error(function(data, status, headers, config) {

     });

    $scope.pie_labels = ["Design", "Development", "Testing"];
    $scope.pie_data = [300, 500, 100];
}]);
