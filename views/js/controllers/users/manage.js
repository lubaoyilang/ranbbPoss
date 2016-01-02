angular.module("app").controller('ManageUserCtrl', ['$scope','DTOptionsBuilder','DTColumnBuilder', function($scope,DTOptionsBuilder,DTColumnBuilder){
	$scope.dataTable = {}
	$scope.dataTable.dtOptions = DTOptionsBuilder.fromSource('data/datatables.json')
		.withDOM('<"table-header"f<t>i>')
        .withScroller()
        .withOption('deferRender', true)
        .withOption('scrollY', 400)
        .withOption('bScrollCollapse', true)
        .withBootstrap()
    $scope.dataTable.dtColumns = [
        DTColumnBuilder.newColumn('id').withTitle('ID'),
        DTColumnBuilder.newColumn('firstName').withTitle('First name'),
        DTColumnBuilder.newColumn('lastName').withTitle('Last name'),
    ];
}]);