<script setup lang="ts">
import { onLaunch, onShow, onHide } from "@dcloudio/uni-app";
onLaunch(() => {
  console.log("App Launch");
});
onShow(() => {
  console.log("App Show");
});
onHide(() => {
  console.log("App Hide");
});
</script>
<style>
	/* #ifndef APP-PLUS-NVUE */
	/* uni.css - 通用组件、模板样式库，可以当作一套ui库应用 */
	@import './assets/css/uni.css';
	/* H5 兼容 pc 所需 */
	/* #ifdef H5 */
	@media screen and (min-width: 768px) {
		body {
			overflow-y: scroll;
		}
	}
	/* 顶栏通栏样式 */
	/* .uni-top-window {
	    left: 0;
	    right: 0;
	} */
	uni-page-body {
		background-color: #F5F5F5 !important;
		min-height: 100% !important;
		height: auto !important;
	}
	.uni-top-window uni-tabbar .uni-tabbar {
		background-color: #fff !important;
	}
	.uni-app--showleftwindow .hideOnPc {
		display: none !important;
	}
	/* #endif */
	/* 以下样式用于 hello uni-app 演示所需 */
	page {
		background-color: #efeff4;
		height: 100%;
		font-size: 28rpx;
		/* line-height: 1.8; */
	}
	.fix-pc-padding {
		padding: 0 50px;
	}
	.uni-header-logo {
		padding: 30rpx;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		margin-top: 10rpx;
	}
	.uni-header-image {
		width: 100px;
		height: 100px;
	}
	.uni-hello-text {
		color: #7A7E83;
	}
	.uni-hello-addfile {
		text-align: center;
		line-height: 300rpx;
		background: #FFF;
		padding: 50rpx;
		margin-top: 10px;
		font-size: 38rpx;
		color: #808080;
	}
	/* #endif*/
</style>

<style>
p {
  font-size: 21px;
  line-height: 1.6em;
  color: #000;
  font-family: '楷体';
}

a:link {
  color: #4e6ef2;
  font-weight: normal;
  text-decoration: none;
}

@keyframes dh {
  0% {
    background-position: 0 0;
  }

  100% {
    background-position: -100% 0;
  }
}

a .a_link_1 {
  color: blue
}

a .a_link_2 {
  color: #20B2AA
}

a .a_link_3 {
  color: #8A2BE2
}

a .a_link_4 {
  color: #FF8C00;
  animation: dh 4s infinite linear
}

a .a_link_5 {
  color: #DAA520;
  animation: dh 4s infinite linear
}

a .a_link_6 {
  background: -webkit-linear-gradient(left, #3498db, #f47920 10%, #d71345 20%, #f7acbc 30%, #ffd400 40%, #3498db 50%, #f47920 60%, #d71345 70%, #f7acbc 80%, #ffd400 90%, #3498db);
  background-clip: text;
  animation: dh 4s infinite linear;
  color: transparent;
  background-size: 200% 100%;
}

.txt_1 {
  color: blue
}

.txt_2 {
  color: #20B2AA
}

.txt_3 {
  color: #8A2BE2
}

.txt_4 {
  color: #FF8C00
}

.txt_5 {
  color: #DAA520;
}

.a_btn {
  outline: none;
  border: none;
  color: #4e6ef2;
  font-weight: normal;
  text-decoration: none;
  font-size: 21px;
  line-height: 1.6em;
  font-family: '楷体';
  background: rgba(0, 0, 0, 0);
  padding: 0;
}

.txt_6 {
  background: linear-gradient(45deg, #DC143C, #FF1493, #00BFFF, #1E90FF);
  background-clip: text;
  color: transparent;
}
</style>