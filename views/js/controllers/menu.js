angular.module("app").controller("MenuCtrl", [
    "$rootScope", "$scope", "$state","$timeout","$http",
    function($rootScope, $scope, $state,$timeout,$http) {


        $scope.logout = function () {
          $http({method: 'POST', url: '/ranbb/logout'})
          .success(function(data, status, headers, config) {
              $state.go("sessions.signin")
          })
          .error(function(data, status, headers, config) {
             $state.go("sessions.signin")
          });
        }

        $rootScope.menu_control = {
            is_expand: false,
            is_active: function(state) {
                if ($state.current.name.indexOf(state) > -1) {
                    return true;
                }
                return false;
            },
            submenu_toggled: function(open){
                $timeout(function() {
                    if(angular.element(document.querySelector("body .expand-dropdown.active, body .expand-dropdown.open")).length > 0) {
                        $rootScope.menu_control.is_expand = true;
                    } else {
                        $rootScope.menu_control.is_expand = false;
                    }
                },200);
            }
        };
        $rootScope.menu_control.submenu_toggled('');
    }
]);
