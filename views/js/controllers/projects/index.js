angular.module("app").controller('ProjectIndexCtrl', ['$scope','$modal','$timeout','$state','$http','ngToast','$rootScope',function($scope,$modal,$timeout,$state,$http,ngToast,$rootScope) {

    $scope.modal = {
        open: function(size,class_name,project) {
            var modalInstance = $modal.open({
                templateUrl: 'modal_example.html',
                controller: 'ShopEditAndAddController',
                size: size,
                windowClass: class_name
            });

            modalInstance.shop  = project;

            modalInstance.result.then(function(result) {
            }, function() {
            });
        }
    }

    $scope.gotoEditShop = function(project){
        $rootScope.shop = project;
        $state.go('main.projects.editShop');
    }

   $scope.project = {
        loading: false,
        page:1,
        size:20,
        scroll: function() {
            if (!$scope.project.loading){
                $scope.project.loading = true;
                $http({method: 'POST',
                    url: '/ranbb/getShopList',
                    params:{page:$scope.project.page,size:$scope.project.size}
                })
                    .success(function(data, status, headers, config) {
                        console.log(data);
                        if(data.Code > 0&& data.Data != null){
                            $scope.project.page++;
                            Array.prototype.push.apply( $scope.projects, data.Data);

                        }else{
                            ngToast.create({
                                content: data.Data,
                                className: 'warning',
                                dismissButton: true
                            });
                        }
                        $scope.project.loading = false;
                    })
                    .error(function(data, status, headers, config) {
                        ngToast.create({
                            content: "加载失败",
                            className: 'warning',
                            dismissButton: true
                        });
                        $scope.project.loading = false;
                    });
            }
        }
    }
    $scope.projects = [];

}]);
