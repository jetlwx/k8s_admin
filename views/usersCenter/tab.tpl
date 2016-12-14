<div>
    <ul class="nav nav-tabs">
         <li role="presentation" {{if .IsuserList}} class="active" {{end}}><a href="/UsersCenter/UserList">User List</a></li>
        <li role="presentation" {{if .Isadduser}} class="active" {{end}}><a href="/UsersCenter/AddUser">Add User</a></li>
        <li role="presentation" {{if .IssetSSHclient}} class="active" {{end}}><a href="/UsersCenter/SSHclient">SSH Client Set</a></li>
    </ul>
</div>
<div>

 
 