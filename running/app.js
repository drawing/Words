
var app = angular.module('remenber', []);

app.controller('RememberController', function($scope, $http) {
	$scope.articles = []
	$scope.familiar_num = 0
	$scope.username = ""
	
	$http({
		method:'GET',
		url:'/articles'
	}).success(
		function(response, status, headers, config) {
			if (response.code == 0) {
				$scope.articles = response.articles;
				$scope.familiar_num = response.familiar_num
				$scope.username = response.username
			} else {
				document.location = response.target
			}
		}
	).error(
		function(response, status, headers, config) {
			alert("获取数据失败")
			$scope.articles = [];
		}
	)
	
	$scope.addfamiliar = new Object();
	$scope.importdict = [];

	$scope.view = new Object();
	$scope.view.title = ""
	
	$scope.review = new Object();
	$scope.review.article = ""
	$scope.review.vocabulary = ["w1", "w2", "w3"]
	
	$scope.action_view_article = function (index, article) {
		if ($scope.articles[index].Name != article) {
			return alert("view error:" + article);
		}
		$scope.view.title = $scope.articles[index].Name;
		
		$scope.view.content = $scope.articles[index].Content
		if ($scope.view.content == null) {
			var view_index = index
			var name = $scope.articles[view_index].Name
			var course_id = $scope.articles[view_index].ID
			$http({
				method:'POST',
				url:'/article_detail',
				params:{"course":name, "id":course_id}
			}).success(
				function(response, status, headers, config) {
					if (response.code == 0) {
						$scope.view.content = response.course.Content
					}
					else {
						$scope.view.content = "获取数据失败"
					}
					
				}
			).error(
				function(response, status, headers, config) {
					$scope.view.content = "获取数据失败"
				}
			)
		}
		$("#page-view-article").modal({keyboard:false, backdrop:"static"});
	};
	
	$scope.action_delete_article = function (index, article) {
		if ($scope.articles[index].Name != article) {
			return alert("delete error:" + article);
		}
		course = $scope.articles[index].Name
		id = $scope.articles[index].ID
		$http({
			method:'POST',
			url:'/delete_article',
			params:{"course":course, "id":id}
		}).success(
			function(response, status, headers, config) {
				if (response.code != 0) {
					alert("删除文章失败")
					return
				}
				$scope.articles.splice(index, 1);
			}
		).error(
			function(response, status, headers, config) {
				alert("获取数据失败")
				return
			}
		)
	};
	
	$scope.action_add_article = function (new_article) {
		// $scope.articles.push(new_article.title);
		
		article = new Object()
		article.ID = "0"
		article.Name = new_article.title
		article.Content = new_article.content
		
		$http({
			method:'POST',
			url:'/add_article',
			headers: {'Content-Type': 'application/x-www-form-urlencoded'},
			transformRequest: function(obj) {
		        var str = [];
		        for(var p in obj)
		        str.push(encodeURIComponent(p) + "=" + encodeURIComponent(obj[p]));
		        return str.join("&");
		    },
			data:{"title":new_article.title, "content":new_article.content}
		}).success(
			function(response, status, headers, config) {
				if (response.code != 0) {
					alert("添加失败")
					return
				}
				$scope.articles.push(article)
			}
		).error(
			function(response, status, headers, config) {
				alert("获取数据失败")
				return
			}
		)
		$("#page-create-article").modal("hide")
		new_article.title = "";
		new_article.content = "";
	};
	
	$scope.actioin_review_vocabulary = function (index, article) {
		$http({
			method:'POST',
			url:'/vocabulary',
			params:{"course":article}
		}).success(
			function(response, status, headers, config) {
				if (response.code != 0) {
					alert("获取单词失败")
					return
				}
				$scope.review.vocabulary = response.vocabulary
				$scope.review.index = 0
				if ($scope.review.vocabulary == null) {
					$scope.review.word = "本课单词已全部完成"
				}
				else {
					$scope.review.word = $scope.review.vocabulary[$scope.review.index]
					// $scope.action_display_translation($scope.review.word)
					$scope.review.trans = new Object()
				}
			}
		).error(
			function(response, status, headers, config) {
				alert("获取数据失败")
				return
			}
		)
		
		$scope.review.audio = document.createElement('audio');
		$scope.review.title = article
		$scope.review.index = 0
		$scope.review.word = ""
		$("#page-review-vocabulary").modal({keyboard:false, backdrop:"static"});
	};
	
	$scope.actioin_add_familiar = function (content) {
		$http({
			method:'POST',
			url:'/add_familiar',
			headers: {'Content-Type': 'application/x-www-form-urlencoded'},
			transformRequest: function(obj) {
		        var str = [];
		        for(var p in obj)
		        str.push(encodeURIComponent(p) + "=" + encodeURIComponent(obj[p]));
		        return str.join("&");
		    },
			data:{"content":content}
		}).success(
			function(response, status, headers, config) {
				if (response.code != 0) {
					alert("添加词汇失败")
					return
				}
				$("#page-add-familiar").modal("hide");
				document.location.reload()
			}
		).error(
			function(response, status, headers, config) {
				alert("获取数据失败")
				return
			}
		)
	};
	
	$scope.actioin_show_add_article = function () {
		$("#page-create-article").modal({keyboard:false, backdrop:"static"});
	};
	$scope.actioin_show_add_familiar = function () {
		$scope.addfamiliar.content = "";
		$("#page-add-familiar").modal({keyboard:false, backdrop:"static"});
	};
	$scope.actioin_show_import_dict = function () {
		$http({
			method:'POST',
			url:'/dictionary_list',
		}).success(
			function(response, status, headers, config) {
				if (response.code != 0) {
					alert("获取词典失败")
					return
				}
				$scope.importdict = response.dict
			}
		).error(
			function(response, status, headers, config) {
				alert("获取数据失败")
			}
		)
		$("#page-import-dict").modal({keyboard:false, backdrop:"static"});
	};
	
	$scope.actioin_import_dict = function (index, dictname) {
		if ($scope.importdict[index].Name != dictname) {
			alert("导入失败");
			return
		}
		content = $scope.importdict[index].Content.toString()
		$http({
			method:'POST',
			url:'/add_familiar',
			headers: {'Content-Type': 'application/x-www-form-urlencoded'},
			transformRequest: function(obj) {
		        var str = [];
		        for(var p in obj)
		        str.push(encodeURIComponent(p) + "=" + encodeURIComponent(obj[p]));
		        return str.join("&");
		    },
			data:{"content":content}
		}).success(
			function(response, status, headers, config) {
				if (response.code != 0) {
					alert("添加词汇失败")
				}
				else {
					alert("添加词汇成功")
				}
			}
		).error(
			function(response, status, headers, config) {
				alert("获取数据失败")
				return
			}
		)
	};
	
	$scope.actioin_review_next_word = function () {
		$scope.review.index ++;
		$scope.review.trans = new Object()
		if ($scope.review.index >= $scope.review.vocabulary.length) {
			$scope.review.index = 0;
		}
		if ($scope.review.vocabulary.length == 0) {
			$scope.review.word = "本课单词已全部完成"
		}
		else {
			$scope.review.word = $scope.review.vocabulary[$scope.review.index]
			// $scope.action_display_translation($scope.review.word)
		}
	};
	$scope.actioin_review_familiar_word = function () {
		var review_word = $scope.review.word
		$http({
			method:'POST',
			url:'/familiar_word',
			params:{"word":review_word}
		}).success(
			function(response, status, headers, config) {
				if (response.code == 0) {
					$scope.familiar_num += 1
				}
			}
		)
		$scope.review.vocabulary.splice($scope.review.index, 1);
		$scope.review.index --;
		$scope.actioin_review_next_word();
	};
	
	$scope.action_display_translation = function (word) {
		$scope.review.trans = new Object()
		$http({
			method:'POST',
			url:'/translate',
			params:{"word":word},
		}).success(
			function(response, status, headers, config) {
				if (response.code == 0) {
					$scope.review.trans = response.dict
				}
			}
		)
	}
	$scope.action_play_sound = function(sound) {
		$scope.review.audio.src = sound;
		$scope.review.audio.play();
	}
});

