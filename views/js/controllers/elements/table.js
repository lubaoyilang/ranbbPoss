angular.module("app").controller('ElementTableCtrl', ['$scope','DTOptionsBuilder','DTColumnBuilder', function($scope,DTOptionsBuilder,DTColumnBuilder){
	$scope.generalTable = {
		classOptions: [
			'Bordered', 
		    'Hover', 
		    'Striped'
		],
		classSelected: []
	}

	$scope.dataTable = {}
	$scope.dataTable.dtOptions = DTOptionsBuilder.fromSource('data/datatables.json')
		.withDOM('<"table-header"lf<t>ip>')
        .withBootstrap()
        .withBootstrapOptions({
            pagination: {
                classes: {
                    ul: 'pagination pagination-sm'
                }
            }
        });
    $scope.dataTable.dtColumns = [
        DTColumnBuilder.newColumn('id').withTitle('ID'),
        DTColumnBuilder.newColumn('firstName').withTitle('First name'),
        DTColumnBuilder.newColumn('lastName').withTitle('Last name')
    ];

    $scope.dataTableWithScroller = {}
    $scope.dataTableWithScroller.dtOptions = DTOptionsBuilder.fromSource('data/datatables.json')
        .withDOM('<"table-header"f<t>i>')
        .withScroller()
        .withOption('deferRender', true)
        .withOption('scrollY', 400)
        .withOption('bScrollCollapse', true)
        .withBootstrap()
    $scope.dataTableWithScroller.dtColumns = [
        DTColumnBuilder.newColumn('id').withTitle('ID'),
        DTColumnBuilder.newColumn('firstName').withTitle('First name'),
        DTColumnBuilder.newColumn('lastName').withTitle('Last name')
    ];
}]);