if(!self.define){let e,a={};const s=(s,n)=>(s=new URL(s+".js",n).href,a[s]||new Promise((a=>{if("document"in self){const e=document.createElement("script");e.src=s,e.onload=a,document.head.appendChild(e)}else e=s,importScripts(s),a()})).then((()=>{let e=a[s];if(!e)throw new Error(`Module ${s} didn’t register its module`);return e})));self.define=(n,t)=>{const c=e||("document"in self?document.currentScript.src:"")||location.href;if(a[c])return;let i={};const r=e=>s(e,c),d={module:{uri:c},exports:i,require:r};a[c]=Promise.all(n.map((e=>d[e]||r(e)))).then((e=>(t(...e),i)))}}define(["./workbox-8e51679a"],(function(e){"use strict";importScripts(),self.skipWaiting(),e.clientsClaim(),e.precacheAndRoute([{url:"/next/_next/static/chunks/15-905c9df63527c447.js",revision:"905c9df63527c447"},{url:"/next/_next/static/chunks/15-905c9df63527c447.js.map",revision:"b409d190f43665716340088179acfb38"},{url:"/next/_next/static/chunks/203-86e0373505403963.js",revision:"86e0373505403963"},{url:"/next/_next/static/chunks/203-86e0373505403963.js.map",revision:"054b552f25b64ae05552facd7bdaffae"},{url:"/next/_next/static/chunks/2edb282b-8765b8d66bb5c3a4.js",revision:"8765b8d66bb5c3a4"},{url:"/next/_next/static/chunks/2edb282b-8765b8d66bb5c3a4.js.map",revision:"2e5cc1e98a9041be17a0fe2c971361bc"},{url:"/next/_next/static/chunks/354-0a997b5f357b7789.js",revision:"0a997b5f357b7789"},{url:"/next/_next/static/chunks/354-0a997b5f357b7789.js.map",revision:"c4171363f99829631247ad7eaf95098a"},{url:"/next/_next/static/chunks/370-1339c4e6601563dd.js",revision:"1339c4e6601563dd"},{url:"/next/_next/static/chunks/370-1339c4e6601563dd.js.map",revision:"dbcf9a3fb73927b22ef8fcb47d9f55b5"},{url:"/next/_next/static/chunks/432-d600282737f8cf97.js",revision:"d600282737f8cf97"},{url:"/next/_next/static/chunks/432-d600282737f8cf97.js.map",revision:"db5c00dc3e434926ae02502877db675a"},{url:"/next/_next/static/chunks/579-d5a6e64d8b4b47d2.js",revision:"d5a6e64d8b4b47d2"},{url:"/next/_next/static/chunks/579-d5a6e64d8b4b47d2.js.map",revision:"393945e0080152bdf8522748fd9cc928"},{url:"/next/_next/static/chunks/75-d29d26e7da10cb33.js",revision:"d29d26e7da10cb33"},{url:"/next/_next/static/chunks/75-d29d26e7da10cb33.js.map",revision:"2978f2ec1f7a742cc7056ca1d1871591"},{url:"/next/_next/static/chunks/956-0f701a29b83c343e.js",revision:"0f701a29b83c343e"},{url:"/next/_next/static/chunks/956-0f701a29b83c343e.js.map",revision:"6c8b63f026760cc06381ef7185356633"},{url:"/next/_next/static/chunks/97cc2b9f-4a814be3c47b7050.js",revision:"4a814be3c47b7050"},{url:"/next/_next/static/chunks/97cc2b9f-4a814be3c47b7050.js.map",revision:"80e5d83bc74ed287c70063870cc1befe"},{url:"/next/_next/static/chunks/framework-154530caa749721d.js",revision:"154530caa749721d"},{url:"/next/_next/static/chunks/framework-154530caa749721d.js.map",revision:"e7237f050a12c2d48c94783391efc1ca"},{url:"/next/_next/static/chunks/main-845e2a21e0a7eed0.js",revision:"845e2a21e0a7eed0"},{url:"/next/_next/static/chunks/main-845e2a21e0a7eed0.js.map",revision:"b1af4a44f42dda92fb19dd48a1f4f068"},{url:"/next/_next/static/chunks/pages/_app-6008db4d54a96da2.js",revision:"6008db4d54a96da2"},{url:"/next/_next/static/chunks/pages/_error-2a1b91330324eb70.js",revision:"2a1b91330324eb70"},{url:"/next/_next/static/chunks/pages/_error-2a1b91330324eb70.js.map",revision:"6344db34081d280a3573cf5de1843669"},{url:"/next/_next/static/chunks/pages/about-ba916ca6f0a231b3.js",revision:"ba916ca6f0a231b3"},{url:"/next/_next/static/chunks/pages/about-ba916ca6f0a231b3.js.map",revision:"3fc86167590392a9ce137ffd560c01c1"},{url:"/next/_next/static/chunks/pages/admin/groups-479b0c399cbdd56a.js",revision:"479b0c399cbdd56a"},{url:"/next/_next/static/chunks/pages/admin/groups-479b0c399cbdd56a.js.map",revision:"52344b4f9a3e34dd0a7b159cd3fb5346"},{url:"/next/_next/static/chunks/pages/admin/links-4b83a43e8d58ea02.js",revision:"4b83a43e8d58ea02"},{url:"/next/_next/static/chunks/pages/admin/links-4b83a43e8d58ea02.js.map",revision:"aa5aa6f353c892064ca83494ecad7e5a"},{url:"/next/_next/static/chunks/pages/admin/users-5711a1f61430fedd.js",revision:"5711a1f61430fedd"},{url:"/next/_next/static/chunks/pages/admin/users-5711a1f61430fedd.js.map",revision:"6022984c47d9c08d6766d48779381542"},{url:"/next/_next/static/chunks/pages/auth/forgot-00ae6d51a374093e.js",revision:"00ae6d51a374093e"},{url:"/next/_next/static/chunks/pages/auth/forgot-00ae6d51a374093e.js.map",revision:"cfae16175c75fc5ee9df6e65eb6b5a69"},{url:"/next/_next/static/chunks/pages/auth/login-7f6c7aaa5f81db25.js",revision:"7f6c7aaa5f81db25"},{url:"/next/_next/static/chunks/pages/auth/login-7f6c7aaa5f81db25.js.map",revision:"389006acf5ac800d7b58468469cb0a9d"},{url:"/next/_next/static/chunks/pages/auth/registration-14a9a03820cf37a7.js",revision:"14a9a03820cf37a7"},{url:"/next/_next/static/chunks/pages/auth/registration-14a9a03820cf37a7.js.map",revision:"f3681d357a41d11e0811877e694161ca"},{url:"/next/_next/static/chunks/pages/auth/verification-5a5e7dad9beddc1b.js",revision:"5a5e7dad9beddc1b"},{url:"/next/_next/static/chunks/pages/auth/verification-5a5e7dad9beddc1b.js.map",revision:"98b2b3226fc08fd9517808455ee7a150"},{url:"/next/_next/static/chunks/pages/contact-ebc0a150f1796524.js",revision:"ebc0a150f1796524"},{url:"/next/_next/static/chunks/pages/contact-ebc0a150f1796524.js.map",revision:"c3517372e4bcc5aea981b86a22afcd73"},{url:"/next/_next/static/chunks/pages/faq-71a5a3e09dd266fc.js",revision:"71a5a3e09dd266fc"},{url:"/next/_next/static/chunks/pages/faq-71a5a3e09dd266fc.js.map",revision:"5f2b19217fdc32dab8997285255db592"},{url:"/next/_next/static/chunks/pages/index-5cc9dbfd5c60e01e.js",revision:"5cc9dbfd5c60e01e"},{url:"/next/_next/static/chunks/pages/index-5cc9dbfd5c60e01e.js.map",revision:"a1921f4e0baa7677fe79b8b82f11ac28"},{url:"/next/_next/static/chunks/pages/pricing-2f2f04aa4fcb33af.js",revision:"2f2f04aa4fcb33af"},{url:"/next/_next/static/chunks/pages/pricing-2f2f04aa4fcb33af.js.map",revision:"9363a14ad4ef61f282d6443cc512ec65"},{url:"/next/_next/static/chunks/pages/privacy-c57a58a229974092.js",revision:"c57a58a229974092"},{url:"/next/_next/static/chunks/pages/privacy-c57a58a229974092.js.map",revision:"b5d18c1fef34fce324f7b22a11363837"},{url:"/next/_next/static/chunks/pages/user/addUrl-98cca900893da1a2.js",revision:"98cca900893da1a2"},{url:"/next/_next/static/chunks/pages/user/addUrl-98cca900893da1a2.js.map",revision:"5d1245c74ba06dbf822e88550b6f742a"},{url:"/next/_next/static/chunks/pages/user/audit-f669cc10c4bb2552.js",revision:"f669cc10c4bb2552"},{url:"/next/_next/static/chunks/pages/user/audit-f669cc10c4bb2552.js.map",revision:"12e01c907680585e8e17eaacd197e328"},{url:"/next/_next/static/chunks/pages/user/billing-051e47ef84d0709a.js",revision:"051e47ef84d0709a"},{url:"/next/_next/static/chunks/pages/user/billing-051e47ef84d0709a.js.map",revision:"0cb2c918ccd534b059be800a906af101"},{url:"/next/_next/static/chunks/pages/user/dashboard-cf238b640f77eb69.js",revision:"cf238b640f77eb69"},{url:"/next/_next/static/chunks/pages/user/dashboard-cf238b640f77eb69.js.map",revision:"c61f1cad7f274ce156ec200811d58771"},{url:"/next/_next/static/chunks/pages/user/integrations-f06c985bb016e5b4.js",revision:"f06c985bb016e5b4"},{url:"/next/_next/static/chunks/pages/user/integrations-f06c985bb016e5b4.js.map",revision:"77ce435de83afb89525f5bf8fc3cfb90"},{url:"/next/_next/static/chunks/pages/user/links-ad7286fea42b8f5d.js",revision:"ad7286fea42b8f5d"},{url:"/next/_next/static/chunks/pages/user/links-ad7286fea42b8f5d.js.map",revision:"04ce5eac5071d0b73a0eb3d4417f7ef1"},{url:"/next/_next/static/chunks/pages/user/profile-4ecc58bd127e2c0a.js",revision:"4ecc58bd127e2c0a"},{url:"/next/_next/static/chunks/pages/user/profile-4ecc58bd127e2c0a.js.map",revision:"17b0aa92dd5996626bb8afdb42c85b8d"},{url:"/next/_next/static/chunks/pages/user/reports-d2c8f9767942541b.js",revision:"d2c8f9767942541b"},{url:"/next/_next/static/chunks/pages/user/reports-d2c8f9767942541b.js.map",revision:"eb66bcd3c65ebae615fa1d9585615af8"},{url:"/next/_next/static/chunks/polyfills-c67a75d1b6f99dc8.js",revision:"837c0df77fd5009c9e46d446188ecfd0"},{url:"/next/_next/static/chunks/webpack-aad7fd3171e1ed4c.js",revision:"aad7fd3171e1ed4c"},{url:"/next/_next/static/chunks/webpack-aad7fd3171e1ed4c.js.map",revision:"4e44dd127a958d08f06b48afbf3a8556"},{url:"/next/_next/static/css/9ebf271484b1097b.css",revision:"9ebf271484b1097b"},{url:"/next/_next/static/css/9ebf271484b1097b.css.map",revision:"5c275425869fd7974e283a4139dc3769"},{url:"/next/_next/static/oP5r0TyEIc74cN_r_1tBI/_buildManifest.js",revision:"97e02b35082052302013bb7eed182c94"},{url:"/next/_next/static/oP5r0TyEIc74cN_r_1tBI/_ssgManifest.js",revision:"b6652df95db52feb4daf4eca35380933"}],{ignoreURLParametersMatching:[]}),e.cleanupOutdatedCaches(),e.registerRoute("/next",new e.NetworkFirst({cacheName:"start-url",plugins:[{cacheWillUpdate:async({response:e})=>e&&"opaqueredirect"===e.type?new Response(e.body,{status:200,statusText:"OK",headers:e.headers}):e}]}),"GET"),e.registerRoute(/^https:\/\/fonts\.(?:gstatic)\.com\/.*/i,new e.CacheFirst({cacheName:"google-fonts-webfonts",plugins:[new e.ExpirationPlugin({maxEntries:4,maxAgeSeconds:31536e3})]}),"GET"),e.registerRoute(/^https:\/\/fonts\.(?:googleapis)\.com\/.*/i,new e.StaleWhileRevalidate({cacheName:"google-fonts-stylesheets",plugins:[new e.ExpirationPlugin({maxEntries:4,maxAgeSeconds:604800})]}),"GET"),e.registerRoute(/\.(?:eot|otf|ttc|ttf|woff|woff2|font.css)$/i,new e.StaleWhileRevalidate({cacheName:"static-font-assets",plugins:[new e.ExpirationPlugin({maxEntries:4,maxAgeSeconds:604800})]}),"GET"),e.registerRoute(/\.(?:jpg|jpeg|gif|png|svg|ico|webp)$/i,new e.StaleWhileRevalidate({cacheName:"static-image-assets",plugins:[new e.ExpirationPlugin({maxEntries:64,maxAgeSeconds:2592e3})]}),"GET"),e.registerRoute(/\/_next\/static.+\.js$/i,new e.CacheFirst({cacheName:"next-static-js-assets",plugins:[new e.ExpirationPlugin({maxEntries:64,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\/_next\/image\?url=.+$/i,new e.StaleWhileRevalidate({cacheName:"next-image",plugins:[new e.ExpirationPlugin({maxEntries:64,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:mp3|wav|ogg)$/i,new e.CacheFirst({cacheName:"static-audio-assets",plugins:[new e.RangeRequestsPlugin,new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:mp4|webm)$/i,new e.CacheFirst({cacheName:"static-video-assets",plugins:[new e.RangeRequestsPlugin,new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:js)$/i,new e.StaleWhileRevalidate({cacheName:"static-js-assets",plugins:[new e.ExpirationPlugin({maxEntries:48,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:css|less)$/i,new e.StaleWhileRevalidate({cacheName:"static-style-assets",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\/_next\/data\/.+\/.+\.json$/i,new e.StaleWhileRevalidate({cacheName:"next-data",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:json|xml|csv)$/i,new e.NetworkFirst({cacheName:"static-data-assets",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute((({sameOrigin:e,url:{pathname:a}})=>!(!e||a.startsWith("/api/auth/")||!a.startsWith("/api/"))),new e.NetworkFirst({cacheName:"apis",networkTimeoutSeconds:10,plugins:[new e.ExpirationPlugin({maxEntries:16,maxAgeSeconds:86400})]}),"GET"),e.registerRoute((({request:e,url:{pathname:a},sameOrigin:s})=>"1"===e.headers.get("RSC")&&"1"===e.headers.get("Next-Router-Prefetch")&&s&&!a.startsWith("/api/")),new e.NetworkFirst({cacheName:"pages-rsc-prefetch",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute((({request:e,url:{pathname:a},sameOrigin:s})=>"1"===e.headers.get("RSC")&&s&&!a.startsWith("/api/")),new e.NetworkFirst({cacheName:"pages-rsc",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute((({url:{pathname:e},sameOrigin:a})=>a&&!e.startsWith("/api/")),new e.NetworkFirst({cacheName:"pages",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute((({sameOrigin:e})=>!e),new e.NetworkFirst({cacheName:"cross-origin",networkTimeoutSeconds:10,plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:3600})]}),"GET")}));
//# sourceMappingURL=sw.js.map
