/**
 * Created by wangmeng on 16/1/3.
 */
angular.module("app").controller('UserEssayController', ['$scope', '$modal', '$timeout', function($scope, $modal, $timeout) {

    $scope.searchText = "";

    $scope.todo = {
        loading: false,
        scroll: function() {
            if (!$scope.todo.loading) {
                $scope.todo.loading = true;
                $timeout(function() {
                    $scope.todo.list.push({
                        title: "申请取现",
                        date: new Date(),
                        status: "success",
                    }, {
                        title: "陈赫",
                        date: new Date(),
                        status: "danger"
                    }, {
                        title: "李晨",
                        date: new Date(),
                        status: "primary"
                    }, {
                        title: "马英九",
                        date: new Date(),
                        status: "danger"
                    }, {
                        title: "科比.",
                        date: new Date(),
                        status: "danger"
                    }, {
                        title: "乔布斯",
                        date: new Date(),
                        status: "primary"
                    }, {
                        title: "黄渤",
                        date: new Date(),
                        status: "success"
                    }, {
                        title: "奥巴马",
                        date: new Date(),
                        status: "warning"
                    });
                    $scope.todo.loading = false;
                }, 1000);
            }
        },
        list: [ {
            title: "用户李晨接单成功",
            date: new Date(),
            status: "success"
        },{
            title: "用户 陈赫 申请取现",
            date: new Date(),
            status: "success",
        },{
            title: "用户 陈赫 申请取现",
            date: new Date(),
            status: "success",
        },{
            title: "用户 陈赫 申请取现",
            date: new Date(),
            status: "success",
        },{
            title: "用户 陈赫 申请取现",
            date: new Date(),
            status: "success",
        },{
            title: "用户 陈赫 申请取现",
            date: new Date(),
            status: "success",
        },{
            title: "用户 陈赫 申请取现",
            date: new Date(),
            status: "success",
        },{
            title: "用户 陈赫 申请取现",
            date: new Date(),
            status: "success",
        },{
            title: "用户 陈赫 申请取现",
            date: new Date(),
            status: "success",
        },{
            title: "用户 陈赫 申请取现",
            date: new Date(),
            status: "success",
        },{
            title: "用户 陈赫 申请取现",
            date: new Date(),
            status: "success",
        }]
    };
}]);
