{{template "/headCssJs.tpl" .}}

		<div class="container" style="clear:both;">
			<ul id="gn-menu" class="gn-menu-main">
				<li class="gn-trigger">
					<a class="gn-icon gn-icon-menu"><span>Menu</span></a>
					<nav class="gn-menu-wrapper">
						<div class="gn-scroller">
							<ul class="gn-menu" style="bolder-bottom:solid 3px red;">
								<li class="gn-search-item">
                                <form action="/distribution" method="request">
									<input placeholder="镜像查询" type="search" class="gn-search" name="search">
									<a class="gn-icon gn-icon-search" href="/distribution"><span>镜像查询</span></a>
                                  </form>
								</li>
								<li>
									<a class="gn-icon gn-icon-download" href="/distribution">镜像管理</a>
                                    <!--
									<ul class="gn-submenu">
										<li><a class="gn-icon gn-icon-illustrator">Vector Illustrations</a></li>
										<li><a class="gn-icon gn-icon-photoshop">Photoshop files</a></li>
									</ul>
                                    -->
								</li>
                                <li>
									<a class="gn-icon gn-icon-archive" href="/kubecluster/listComponets">集群管理</a>
									<ul class="gn-submenu">
										<li><a class="gn-icon gn-icon-article">Articles</a></li>
										<li><a class="gn-icon gn-icon-pictures">Images</a></li>
										<li><a class="gn-icon gn-icon-videos">Videos</a></li>
									</ul>
								</li>
								<li><a class="gn-icon gn-icon-cog" href="/sysconfig">系统设置</a>
                                <!--
                                <ul class="gn-submenu">
										<li><a class="gn-icon gn-icon-article" >仓库设置</a></li>
										<li><a class="gn-icon gn-icon-pictures">Kube设置</a></li>
										<li><a class="gn-icon gn-icon-videos">Videos</a></li>
									</ul>
                                    --></li>
										<li><a class="gn-icon gn-icon-cog" href="/virtualization/machinelist">Virtualization</a>
                              
                                <ul class="gn-submenu">
										<li><a class="gn-icon gn-icon-article" href="/UsersCenter/UserList">Users Center</a></li>
										<li><a class="gn-icon gn-icon-pictures">Kube设置</a></li>
										<li><a class="gn-icon gn-icon-videos">Videos</a></li>
									</ul>
                                   </li>
								<li><a class="gn-icon gn-icon-help">Help</a></li>
								
							</ul>
						</div><!-- /gn-scroller -->
					</nav>
				</li>
				<li><a href="#">Codrops</a></li>
				{{if .Islogin}}
				<li><a class="codrops-icon codrops-icon-prev" ><span>{{.Name}}</span></a></li>
				<li><a class="codrops-icon codrops-icon-prev" href="/login?logout=1"><span>Logout</span></a></li>
				{{else}}
				<li><a class="codrops-icon codrops-icon-prev" href="/login"><span>Login</span></a></li>
				{{end}}
				
				<li><a class="codrops-icon codrops-icon-prev" href="#"><span>Previous Demo</span></a></li>
                
			</ul>
			<div class="jetosmain">
            <div class="maininner">

            <div class="mainbody">