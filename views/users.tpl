<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>User list</title>
</head>
<body>
	<table>
		<tr>
			<th>No</t>
			<th>ID</th>
			<th>Name</th>
		</tr>
		{{range $key, $val := .users}}
		<tr>
			<td>{{ $key }}</td>
			<td>{{$val.Id}}</td>
			<td>{{$val.Name}}</td>
		</tr>
		{{end}}
	</table>
</body>
</html>