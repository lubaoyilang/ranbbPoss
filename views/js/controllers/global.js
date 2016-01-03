angular.module("app").controller("GlobalCtrl", [
    "$rootScope", "$scope", "$state","$timeout",
    function($rootScope, $scope, $state,$timeout) {

        $rootScope.page_control = {
            mode: function() {
                if($state.current.name.indexOf("sessions") > -1) {
                    $state.go("sessions.signin");
                    return "login"
                }
                return ""
            }
        };
    }
]);
