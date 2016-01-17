angular.module("app").controller('AdminListController', ['$scope', '$modal', '$timeout','$http','ngToast',function($scope, $modal, $timeout,$http,ngToast) {

    $scope.searchText = "";

    //获取最新店铺


    $scope.todo = {
        loading: false,
        scroll: function() {
            if (!$scope.todo.loading) {
                $scope.todo.loading = true;

                $http({method: 'POST',
                    url: '/ranbb/getAdminList',
                    params:{'page':$scope.todo.curentPage,'size':10}
                })
                    .success(function(data, status, headers, config) {
                        console.log(data);
                        if(data==null){
                            $scope.todo.loading = false;
                            return
                        }
                        Array.prototype.push.apply( $scope.todo.list, data);
                        $scope.todo.curentPage += 1;
                        $scope.todo.loading = false;
                    })
                    .error(function(data, status, headers, config) {
                        $scope.todo.loading = false;
                    });
            }
        },
        list: [],
        curentPage:1
    };

    $scope.deleteAdmin=function(item){
        $http({method: 'POST',
            url: '/ranbb/deleteAdmin',
            params:{'mobile':item.Mobile}
        })
            .success(function(data, status, headers, config) {
                if(data==null){
                    ngToast.create({
                        content: "删除失败",
                        className: 'warning',
                        dismissButton: true
                    });
                }
                if(data.Code < 0){
                    ngToast.create({
                        content: data.Data,
                        className: 'warning',
                        dismissButton: true
                    });
                    return
                }
                ngToast.create({
                    content: "删除成功",
                    className: 'success',
                    dismissButton: true
                });
                var index = $scope.todo.list.indexOf(item)
                $scope.todo.list.splice(index,1);
                return;
            })
            .error(function(data, status, headers, config) {
                ngToast.create({
                    content: "删除失败",
                    className: 'warning',
                    dismissButton: true
                });
                return;
            });
    };


    $scope.modal = {
        open: function(size,class_name,item) {
            var modalInstance = $modal.open({
                templateUrl: 'admin_edit_modal.html',
                controller: 'adminEditController',
                size: size,
                windowClass: class_name
            });

            modalInstance.user = item;

            modalInstance.result.then(function(admin) {
                $http({method: 'POST',
                    url: '/ranbb/updateAdmin',
                    params:admin
                })
                    .success(function(data, status, headers, config) {
                        if(data==null){
                            ngToast.create({
                                content: "更新失败",
                                className: 'warning',
                                dismissButton: true
                            });
                        }
                        if(data.Code < 0){
                            ngToast.create({
                                content: data.Data,
                                className: 'warning',
                                dismissButton: true
                            });
                            return
                        }
                        ngToast.create({
                            content: "更新成功",
                            className: 'success',
                            dismissButton: true
                        });
                    })
                    .error(function(data, status, headers, config) {
                        ngToast.create({
                            content: "更新失败",
                            className: 'warning',
                            dismissButton: true
                        });
                    });
            }, function() {
                alert("asdf");
            });
        }
    };
}]);