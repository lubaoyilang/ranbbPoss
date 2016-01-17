angular.module("app").controller("GlobalCtrl", [
    "$rootScope", "$scope", "$state","$timeout","$http",
    function($rootScope, $scope, $state,$timeout,$http) {

        $http({method: 'POST', url: '/ranbb/checkSession'})
        .success(function(data, status, headers, config) {
            if (data.Name == null) {
                $state.go('sessions.signin');
            }
        })
        .error(function(data, status, headers, config) {
            $state.go('sessions.signin');
        });

        $rootScope.page_control = {
            mode: function() {
                if($state.current.name.indexOf("sessions") > -1) {
                    return "login"
                }
                return ""
            }
        };
    }
]);
