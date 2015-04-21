app = angular.module("myapp", [])
app.controller "MainController", ["$scope", "$http", ($scope, $http)->
	$scope.page = 0
	$scope.index = 0
	$scope.joke = {msg: "加载中", id: 0}
	# 下一个笑话
	$scope.nextPage = ()->
		$scope.index += 1
		if $scope.jokes.length <= $scope.index 
			$scope.page += 1
			getData($scope, $http)	
		else 
			$scope.joke = $scope.jokes[$scope.index]
	# 上一个笑话
	$scope.prePage = ()->
		$scope.index -= 1
		if $scope.index < 0 
			return $scope.index = 0 if $scope.page <= 0
			$scope.page -= 1
			getData($scope, $http)
		else
			$scope.joke = $scope.jokes[$scope.index]
	# 回复
	$scope.replyBtn = ()->
		data = {id: $scope.joke.id, msg: $scope.newReplyMsg}
		$scope.newReplyMsg = ""
		$scope.joke.replies.push(data)
		$.post("/joke/reply", data)

	getData($scope, $http)
	console.log "init MainController"
]
getData = ($scope, $http)-> 
	$http.get("/joke?page=#{$scope.page}").success((data)-> 
		$scope.jokes = data
		$scope.index = 0
		$scope.joke = $scope.jokes[$scope.index]
	)
