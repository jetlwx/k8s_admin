<div>
    <ul class="nav nav-tabs">
        <li role="presentation" {{if .IsComponets}} class="active" {{end}}><a href="/kubecluster/listComponets">Componets</a></li>
        <li role="presentation" {{if .IsClusternodes}} class="active" {{end}}><a href="/kubecluster/listNodes">Nodes</a></li>
        <li role="presentation" {{if .IsEndpoints}} class="active" {{end}}><a href="/kubecluster/Endpoints">Endpoints</a></li>
        <li role="presentation" {{if .IsServcie}} class="active" {{end}}><a href="/kubecluster/listServices">Service</a></li>
        <li role="presentation" {{if .IsReplication}} class="active" {{end}}><a href="/kubecluster/listReplications">Replication</a></li>
        
        <li role="presentation" {{if .IsPods}} class="active" {{end}}><a href="/kubecluster/listPods">Pods</a></li>
    </ul>
</div>