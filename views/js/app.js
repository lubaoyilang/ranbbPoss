var app;

app = angular.module("app", ["ngSanitize","ngAnimate","ui.router", "ui.bootstrap","ui.bootstrap.datetimepicker","ui.select","checklist-model","datatables","datatables.bootstrap","datatables.scroller","ngToast","infinite-scroll","chart.js","ng-sidr"]).run([
  "$rootScope", "$state", "$stateParams", function($rootScope, $state, $stateParams) {

    //  Config Chart.JS Colors
    Chart.defaults.global.colours[0] = {
      fillColor: "rgba(26,188,156,0.65)",
      pointColor: "rgba(26,188,156,1)",
      pointHighlightFill: "#fff",
      pointHighlightStroke: "rgba(26,188,156,0.8)",
      pointStrokeColor: "#fff",
      strokeColor: "rgba(26,188,156,1)"
    };
    Chart.defaults.global.colours[1] = {
      fillColor: "rgba(96,100,109,0.65)",
      pointColor: "rgba(96,100,109,1)",
      pointHighlightFill: "#fff",
      pointHighlightStroke: "rgba(96,100,109,0.8)",
      pointStrokeColor: "#fff",
      strokeColor: "rgba(96,100,109,1)"
    }
    Chart.defaults.global.colours[2] = {
      fillColor: "rgba(52,152,219,0.65)",
      pointColor: "rgba(52,152,219,1)",
      pointHighlightFill: "#fff",
      pointHighlightStroke: "rgba(52,152,219,0.8)",
      pointStrokeColor: "#fff",
      strokeColor: "rgba(52,152,219,1)"
    }

    $rootScope.$state = $state;
    return $rootScope.$stateParams = $stateParams;
  }
]).config(function($stateProvider, $urlRouterProvider) {
  $urlRouterProvider.otherwise('/main/dashboard');
  $stateProvider.state('main', {
  	abstract: true,
    url: '/main',
    templateUrl: 'templates/layout/main.html'
  }).state('main.dashboard', {
    url: '/dashboard',
    templateUrl: 'templates/controllers/dashboard.html',
    controller: 'DashboardCtrl'
  })
  .state('main.users', {
    abstract: true,
    url: '/users',
    templateUrl: 'templates/controllers/users/main.html'
  })
  .state('main.users.index', {
    url: '/index',
    templateUrl: 'templates/controllers/users/index.html',
    controller: 'ManageUserCtrl'
  })
  .state('main.users.adminList',{
        url:'/adminList',
        templateUrl:'templates/controllers/users/adminList.html',
        controller:'AdminListController'
  })
  .state('main.projects', {
    abstract: true,
    url: '/projects',
    templateUrl: 'templates/controllers/users/main.html'
  })
  .state('main.projects.index', {
    url: '/index',
    templateUrl: 'templates/controllers/projects/index.html',
    controller: 'ProjectIndexCtrl'
  })
  .state('main.projects.goodsList', {
    url: '/goodsList',
    templateUrl: 'templates/controllers/projects/goodsList.html',
    controller: 'GoodsListCtrl'
  })
  .state('main.projects.editShop', {
    url: '/editShop',
    templateUrl: 'templates/controllers/projects/EditShop.html',
    controller: 'EditShopController'
  })
  .state('main.projects.addShop', {
    url: '/editShop',
    templateUrl: 'templates/controllers/projects/addShop.html',
    controller: 'AddShopController'
  })
  .state('main.projects.userEssay', {
    url: '/userEssay',
    templateUrl: 'templates/controllers/projects/userEssay.html',
    controller: 'UserEssayController'
  })
  .state('main.users.create', {
    url: '/create',
    templateUrl: 'templates/controllers/users/create.html',
    controller:'CreateAdminController'
  })
  .state('main.example', {
    abstract: true,
    url: '/example',
    templateUrl: 'templates/controllers/example/main.html'
  })
  .state('main.example.timeline', {
    url: '/timeline',
    templateUrl: 'templates/controllers/example/timeline.html',
    controller: "ExampleTimeLineCtrl"
  })
  .state('main.elements', {
    abstract: true,
    url: '/elements',
    templateUrl: 'templates/controllers/elements/main.html'
  })
  .state('main.elements.index', {
    url: '/index',
    templateUrl: 'templates/controllers/elements/index.html'
  })
  .state('main.elements.components', {
    url: '/components',
    templateUrl: 'templates/controllers/elements/components.html',
    controller: "ElementComponentsCtrl"
  })
  .state('main.elements.form', {
    url: '/form',
    templateUrl: 'templates/controllers/elements/form.html'
  })
  .state('main.elements.table', {
    url: '/table',
    templateUrl: 'templates/controllers/elements/table.html',
    controller: 'ElementTableCtrl'
  })
  .state('main.elements.chartjs', {
    url: '/chartjs',
    templateUrl: 'templates/controllers/elements/chartjs.html',
    controller: 'ElementChartJSCtrl'
  })
  .state('main.elements.infinite_scroll_table', {
    url: '/infinite_scroll_table',
    templateUrl: 'templates/controllers/elements/infinite_scroll_table.html',
    controller: 'ElementInfiniteScrollTableCtrl'
  })
  .state('main.elements.infinite_scroll_grid', {
    url: '/infinite_scroll_grid',
    templateUrl: 'templates/controllers/elements/infinite_scroll_grid.html',
    controller: 'ElementInfiniteScrollGridCtrl'
  })
  .state('sessions', {
    url: '/sessions',
    abstract: true,
    templateUrl: 'templates/layout/sessions.html'
  })
  .state('sessions.signin', {
    url: '/signin',
    templateUrl: 'templates/controllers/sessions/signin.html',
    controller: 'SessionSigninCtrl'
  })
  .state('sessions.register', {
    url: '/register',
    templateUrl: 'templates/controllers/sessions/register.html',
    controller: 'SessionRegisterCtrl'
  });
});
