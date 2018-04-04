<!doctype html>
<html lang="en">
  <head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <link rel="stylesheet" href="/static/bootstrap.min.css">
  <title>{{.Title}}</title>
</head>
<body>
<h1>{{.Title}}</h1>

<p>[<a href="/edit/{{.Title}}">edit</a>]</p>

    <div class="container-fluid">
        <div class="row">
            <div class="col-md-8 col-md-offset-1">
                <div class="page-header">
                    <h1>List of <small>something</small></h1>
                </div>
 		<script type="text/javascript">
			function openallurls() {
                window.open("http://link1.com/123/456", "http://link1.com/123/456")
                window.open("http://link2.com/123/456", "http://link2.com/123/456")
			}
		</script>
		<a href="javascript:void(0);" onclick="openallurls()" class="btn btn-primary">Open all links</a>
                <table class="table table-hover table-condensed">
                    <thead>
                    <tr>
                        <th colspan="4">List</th>
                    </tr>
                    </thead>
                    <tbody>
                    <tr>
                        <td>Description</td>
                        <td>
                            <div class="btn-group">
                                <button type="button" class="btn btn-default">-</button>
                                <button type="button" class="btn btn-info counter">5<small>(10)</small></button>
                                <button type="button" class="btn btn-default">+</button>
                            </div>
                        </td>
                        <td><img src="/static/image.jpeg" class="img-responsive" alt="Rutracker"></td>
                        <td><a href="http://link1.com/123/456"><button type="button" class="btn btn-default">&gt;</button></a></td>
                    </tr>
                    </tbody>
                </table>
		<a href="javascript:void(0);" onclick="openallurls()" class="btn btn-primary">Open all links</a>
            </div>
        </div>
    </div>
    <!-- Optional JavaScript -->
    <!-- jQuery first, then Popper.js, then Bootstrap JS -->
    <script src="/static/jquery-3.2.1.slim.min.js"></script>
    <script src="/static/popper.min.js"></script>
    <script src="/static/bootstrap.min.js"></script>
</body>
</html>
