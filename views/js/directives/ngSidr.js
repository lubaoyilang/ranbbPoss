angular.module('ng-sidr', [])
    .directive('ngSidr', function($compile) {
        return {
            restrict: "EAC",
            scope: {
                scope: "@"
            },
            link: function(scope, elm, attrs) {
                elm.sidr({
                    name: 'sidr-main',
                    source: '.' + attrs.nav,
                    side: 'right'
                });

                $(".sidr-inner > li:not(.sidr-class-expand-dropdown) > a, .sidr-inner > li ul li a, .sidr-inner .sidr-class-mobile-menu-header").bind("click", function() {
                    $.sidr('close', 'sidr-main');
                });
            }
        };
    });
