<div>
    <ul class="nav nav-tabs">
     <li role="presentation" {{if .Ismachinelist}} class="active" {{end}}><a href="/virtualization/machinelist" style="font-weight:bold;">Machine List</a></li>
        <li role="presentation" {{if .IsIpPool}} class="active" {{end}}><a href="/virtualization/setting" style="font-weight:bold;">Ip Pool</a></li>
         <li role="presentation" {{if .IsVirtualmachineMother}} class="active" {{end}}><a href="/VirtualHostMOther/Host" style="font-weight:bold;">Hosts</a></li>     
    </ul>
</div>
<div>
