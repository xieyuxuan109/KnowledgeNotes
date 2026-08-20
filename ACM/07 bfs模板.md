## bfs模板
```
#include <iostream>
#include <cstring>
#define N 410
#include <queue>
using namespace std;
typedef pair<int,int> PII;
int m,n,x,y;
int dist[N][N];

int dx[]={1,2,2,1,-1,-2,-2,-1};
int dy[]={2,1,-1,-2,-2,-1,1,2};

void bfs(){
	memset(dist,-1,sizeof dist);
	
	queue<PII> q;
	q.push({x,y});
	dist[x][y]=0;
	
	while(!q.empty()){
		PII t=q.front();
		q.pop();
		for(int i=0;i<8;i++){
			int x1=t.first+dx[i];
			int y1=t.second+dy[i];
			if(x1<1||x1>n||y1<1||y1>m||dist[x1][y1]!=-1)continue;
			dist[x1][y1]=dist[t.first][t.second]+1;
			q.push({x1,y1});
		}
	}
}
int main() {
	cin>>n>>m>>x>>y;
	bfs();
	for(int i=1;i<=n;i++){
		for(int j=1;j<=m;j++){
			cout<<dist[i][j]<<" ";
		}
		cout<<endl;
	}
	return 0;
}
```