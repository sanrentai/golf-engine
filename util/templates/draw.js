var a=document.getElementById("golfcanvas"),b=a.getContext("2d"),c=b.createImageData(768,768),d=9216,e=4608,f=new Uint8Array(1+d+e),g=[0,1,2,3,4,5,6,7];a.addEventListener('contextmenu',e=>e.preventDefault());function h(x,y,r){s=x*4+y*3072;for(let i=s;i<s+4;i++){for(let j=0;j<4;j++){c.data[i*4+j*3072]=r[0];c.data[i*4+1+j*3072]=r[1];c.data[i*4+2+j*3072]=r[2];c.data[i*4+3+j*3072]=255;}}}function drawScreen(){a=i[(f[d+e]&240)/16];b=i[(f[d+e]&15)];for(let i=0;i<4;i++){g[i]=a[i];g[i+4]=b[i];}for(let i=0;i<36864;i++){let j=Math.floor(Math.floor(i/4)/2*3),k=j+(2-j%3),l=(i%4)*2,m=i%8,c=(3<<l&f[j])>>l,p=(1<<m&f[k])>>m,n=g[c+p*4],x=Math.floor(i%192),y=Math.floor(i/192);h(x,y,n)}b.putImageData(c,0,0)}let i=[[[0,0,0],[224,60,40],[255,255,255],[215,215,215]],[[168,168,168],[123,123,123],[52,52,52],[21,21,21]],[[13,32,48],[65,93,102],[113,166,161],[189,255,202]],[[37,226,205],[10,152,172],[0,82,128],[0,96,75]],[[32,181,98],[88,211,50],[19,157,8],[0,78,0]],[[23,40,8],[55,109,3],[106,180,23],[140,214,18]],[[190,235,113],[238,255,169],[182,193,33],[147,151,23]],[[204,143,21],[255,187,49],[255,231,55],[246,143,55]],[[173,78,26],[35,23,18],[92,60,13],[174,108,55]],[[197,151,130],[226,215,181],[79,21,7],[130,60,61]],[[218,101,94],[225,130,137],[245,183,132],[255,233,197]],[[255,130,206],[207,60,113],[135,22,70],[163,40,179]],[[204,105,228],[213,156,252],[254,201,237],[226,201,255]],[[166,117,254],[106,49,202],[90,25,145],[33,22,64]],[[61,52,165],[98,100,220],[155,160,239],[152,220,255]],[[91,168,255],[10,137,255],[2,74,202],[0,23,125]],]