if(!self.define){let e,s={};const a=(a,c)=>(a=new URL(a+".js",c).href,s[a]||new Promise((s=>{if("document"in self){const e=document.createElement("script");e.src=a,e.onload=s,document.head.appendChild(e)}else e=a,importScripts(a),s()})).then((()=>{let e=s[a];if(!e)throw new Error(`Module ${a} didn’t register its module`);return e})));self.define=(c,n)=>{const t=e||("document"in self?document.currentScript.src:"")||location.href;if(s[t])return;let i={};const d=e=>a(e,t),r={module:{uri:t},exports:i,require:d};s[t]=Promise.all(c.map((e=>r[e]||d(e)))).then((e=>(n(...e),i)))}}define(["./workbox-8e51679a"],(function(e){"use strict";importScripts(),self.skipWaiting(),e.clientsClaim(),e.precacheAndRoute([{url:"/next/_next/static/2sZ_LoYkGfAHT2PLoNGyj/_buildManifest.js",revision:"d263cf978ac39cbf3a429c7bfff6bb3d"},{url:"/next/_next/static/2sZ_LoYkGfAHT2PLoNGyj/_ssgManifest.js",revision:"b6652df95db52feb4daf4eca35380933"},{url:"/next/_next/static/chunks/124-9150e62d8145167f.js",revision:"9150e62d8145167f"},{url:"/next/_next/static/chunks/124-9150e62d8145167f.js.map",revision:"bdc5d2cb650cda361d71980a5ada0361"},{url:"/next/_next/static/chunks/15-fedeb987a0cd8bcc.js",revision:"fedeb987a0cd8bcc"},{url:"/next/_next/static/chunks/15-fedeb987a0cd8bcc.js.map",revision:"e8539409629891c83801b261cb8d2a57"},{url:"/next/_next/static/chunks/2edb282b-775571175466ab69.js",revision:"775571175466ab69"},{url:"/next/_next/static/chunks/2edb282b-775571175466ab69.js.map",revision:"e7476d1c47bc1a0de2f9bc932ee870ad"},{url:"/next/_next/static/chunks/354-629d6c78f9fcca32.js",revision:"629d6c78f9fcca32"},{url:"/next/_next/static/chunks/354-629d6c78f9fcca32.js.map",revision:"6de232c5198e0ff9277d3b5c9fa14cf6"},{url:"/next/_next/static/chunks/579-5555697d2acb8489.js",revision:"5555697d2acb8489"},{url:"/next/_next/static/chunks/579-5555697d2acb8489.js.map",revision:"c112c3f4dc7217a5e730083272fad569"},{url:"/next/_next/static/chunks/74-c253286f2c86e5bc.js",revision:"c253286f2c86e5bc"},{url:"/next/_next/static/chunks/74-c253286f2c86e5bc.js.map",revision:"3e3b2b5744b3cef0ea304a0c5c6134b1"},{url:"/next/_next/static/chunks/75-08a71803941eff5d.js",revision:"08a71803941eff5d"},{url:"/next/_next/static/chunks/75-08a71803941eff5d.js.map",revision:"c0410625306319ecbc3d2ef1ebffbf7c"},{url:"/next/_next/static/chunks/927-5e96edc0f71bc1cf.js",revision:"5e96edc0f71bc1cf"},{url:"/next/_next/static/chunks/927-5e96edc0f71bc1cf.js.map",revision:"b9509369aa8489f44aa54d67fc2f11b9"},{url:"/next/_next/static/chunks/97cc2b9f-92ab705079d625be.js",revision:"92ab705079d625be"},{url:"/next/_next/static/chunks/97cc2b9f-92ab705079d625be.js.map",revision:"ec7fe853dd3307219e4e6d053411043e"},{url:"/next/_next/static/chunks/framework-d6095a5336d3e425.js",revision:"d6095a5336d3e425"},{url:"/next/_next/static/chunks/framework-d6095a5336d3e425.js.map",revision:"d2bf73d63d22cb42ec2075cc76b211fd"},{url:"/next/_next/static/chunks/main-d8b8b7f2cfb88da7.js",revision:"d8b8b7f2cfb88da7"},{url:"/next/_next/static/chunks/main-d8b8b7f2cfb88da7.js.map",revision:"1b410046dae429d2f2b38e55b8a06f07"},{url:"/next/_next/static/chunks/pages/_app-dbb54d73515c6120.js",revision:"dbb54d73515c6120"},{url:"/next/_next/static/chunks/pages/_error-f84583a6d713f3ce.js",revision:"f84583a6d713f3ce"},{url:"/next/_next/static/chunks/pages/_error-f84583a6d713f3ce.js.map",revision:"c7979365e737ffad2cff10d3cabc4f78"},{url:"/next/_next/static/chunks/pages/about-f92a563882be0fc5.js",revision:"f92a563882be0fc5"},{url:"/next/_next/static/chunks/pages/about-f92a563882be0fc5.js.map",revision:"240c1e455c81567387bee79569c085d0"},{url:"/next/_next/static/chunks/pages/admin/groups-4ff55b9ff3a23cd9.js",revision:"4ff55b9ff3a23cd9"},{url:"/next/_next/static/chunks/pages/admin/groups-4ff55b9ff3a23cd9.js.map",revision:"693ac2b7f4ac12f11e5299cfc9a7ef44"},{url:"/next/_next/static/chunks/pages/admin/links-1d3d1d455e1098e6.js",revision:"1d3d1d455e1098e6"},{url:"/next/_next/static/chunks/pages/admin/links-1d3d1d455e1098e6.js.map",revision:"f063db13b11e7cb6a4cd84f3c128cb3e"},{url:"/next/_next/static/chunks/pages/admin/users-4735e55853f30e96.js",revision:"4735e55853f30e96"},{url:"/next/_next/static/chunks/pages/admin/users-4735e55853f30e96.js.map",revision:"9526a86cd8edb02112ba4e8fcd7fe4f6"},{url:"/next/_next/static/chunks/pages/auth/forgot-97c6f59a110662ed.js",revision:"97c6f59a110662ed"},{url:"/next/_next/static/chunks/pages/auth/forgot-97c6f59a110662ed.js.map",revision:"62c0e923c19e14ac3239ea346c2e1877"},{url:"/next/_next/static/chunks/pages/auth/login-27bdc42224cb397f.js",revision:"27bdc42224cb397f"},{url:"/next/_next/static/chunks/pages/auth/login-27bdc42224cb397f.js.map",revision:"94bd52fa052f4ea321b5dd61cece727a"},{url:"/next/_next/static/chunks/pages/auth/registration-2a1179e14f22cb5a.js",revision:"2a1179e14f22cb5a"},{url:"/next/_next/static/chunks/pages/auth/registration-2a1179e14f22cb5a.js.map",revision:"84882ad999da76bcf3c0aa8e410d12d7"},{url:"/next/_next/static/chunks/pages/auth/verification-5c33d8c8981a7991.js",revision:"5c33d8c8981a7991"},{url:"/next/_next/static/chunks/pages/auth/verification-5c33d8c8981a7991.js.map",revision:"3209dd1070e715695bdd627387798e91"},{url:"/next/_next/static/chunks/pages/contact-be5da8b4276c3d70.js",revision:"be5da8b4276c3d70"},{url:"/next/_next/static/chunks/pages/contact-be5da8b4276c3d70.js.map",revision:"f967fd677968677ed67fe17cec6737dd"},{url:"/next/_next/static/chunks/pages/faq-e4c6a98922c4200a.js",revision:"e4c6a98922c4200a"},{url:"/next/_next/static/chunks/pages/faq-e4c6a98922c4200a.js.map",revision:"a115cbf87abef144d7daffa8a6dec4e8"},{url:"/next/_next/static/chunks/pages/index-c4a7d8059549fa08.js",revision:"c4a7d8059549fa08"},{url:"/next/_next/static/chunks/pages/index-c4a7d8059549fa08.js.map",revision:"893c791bc0d613390a520eaabacb2bab"},{url:"/next/_next/static/chunks/pages/pricing-34816d3a344b2607.js",revision:"34816d3a344b2607"},{url:"/next/_next/static/chunks/pages/pricing-34816d3a344b2607.js.map",revision:"cfb907fdf608c8e465a2dbcda67c799d"},{url:"/next/_next/static/chunks/pages/privacy-04542a695f63f69d.js",revision:"04542a695f63f69d"},{url:"/next/_next/static/chunks/pages/privacy-04542a695f63f69d.js.map",revision:"6c59dadd8a259f84e9534bfbf41d13f6"},{url:"/next/_next/static/chunks/pages/user/addUrl-db6b9cbd49fd524c.js",revision:"db6b9cbd49fd524c"},{url:"/next/_next/static/chunks/pages/user/addUrl-db6b9cbd49fd524c.js.map",revision:"dc3b4043533085ef0588e7aaa4c4c673"},{url:"/next/_next/static/chunks/pages/user/audit-9141f0bd0f78ccd2.js",revision:"9141f0bd0f78ccd2"},{url:"/next/_next/static/chunks/pages/user/audit-9141f0bd0f78ccd2.js.map",revision:"55edd95d4d125825515014ddc92d0996"},{url:"/next/_next/static/chunks/pages/user/billing-9ece22a20d3cc2f2.js",revision:"9ece22a20d3cc2f2"},{url:"/next/_next/static/chunks/pages/user/billing-9ece22a20d3cc2f2.js.map",revision:"ebf6f85da659dadc81e551620a19bed4"},{url:"/next/_next/static/chunks/pages/user/dashboard-b7f289ea6997ed81.js",revision:"b7f289ea6997ed81"},{url:"/next/_next/static/chunks/pages/user/dashboard-b7f289ea6997ed81.js.map",revision:"d0f9592e6e65a72bd644cee1181f3b93"},{url:"/next/_next/static/chunks/pages/user/integrations-aeaae8a5ad7c9054.js",revision:"aeaae8a5ad7c9054"},{url:"/next/_next/static/chunks/pages/user/integrations-aeaae8a5ad7c9054.js.map",revision:"f8656f7cf1e80bebdfe757540ebd3cd0"},{url:"/next/_next/static/chunks/pages/user/links-84b4fe7082169ca7.js",revision:"84b4fe7082169ca7"},{url:"/next/_next/static/chunks/pages/user/links-84b4fe7082169ca7.js.map",revision:"6fa767a3fdca7894de189af547c1a9fd"},{url:"/next/_next/static/chunks/pages/user/profile-96ed66946f251674.js",revision:"96ed66946f251674"},{url:"/next/_next/static/chunks/pages/user/profile-96ed66946f251674.js.map",revision:"d3f360a4c10f4971759b4fc4b54cec4f"},{url:"/next/_next/static/chunks/pages/user/reports-bd7065096d780683.js",revision:"bd7065096d780683"},{url:"/next/_next/static/chunks/pages/user/reports-bd7065096d780683.js.map",revision:"81ce89ff807d5d84685cfaab18201dfb"},{url:"/next/_next/static/chunks/polyfills-c67a75d1b6f99dc8.js",revision:"837c0df77fd5009c9e46d446188ecfd0"},{url:"/next/_next/static/chunks/webpack-aad7fd3171e1ed4c.js",revision:"aad7fd3171e1ed4c"},{url:"/next/_next/static/chunks/webpack-aad7fd3171e1ed4c.js.map",revision:"4e44dd127a958d08f06b48afbf3a8556"},{url:"/next/_next/static/css/07df0638c2c5e066.css",revision:"07df0638c2c5e066"},{url:"/next/_next/static/css/07df0638c2c5e066.css.map",revision:"9f85a04bc45de40b7d68a3fe29f42a16"}],{ignoreURLParametersMatching:[]}),e.cleanupOutdatedCaches(),e.registerRoute("/next",new e.NetworkFirst({cacheName:"start-url",plugins:[{cacheWillUpdate:async({response:e})=>e&&"opaqueredirect"===e.type?new Response(e.body,{status:200,statusText:"OK",headers:e.headers}):e}]}),"GET"),e.registerRoute(/^https:\/\/fonts\.(?:gstatic)\.com\/.*/i,new e.CacheFirst({cacheName:"google-fonts-webfonts",plugins:[new e.ExpirationPlugin({maxEntries:4,maxAgeSeconds:31536e3})]}),"GET"),e.registerRoute(/^https:\/\/fonts\.(?:googleapis)\.com\/.*/i,new e.StaleWhileRevalidate({cacheName:"google-fonts-stylesheets",plugins:[new e.ExpirationPlugin({maxEntries:4,maxAgeSeconds:604800})]}),"GET"),e.registerRoute(/\.(?:eot|otf|ttc|ttf|woff|woff2|font.css)$/i,new e.StaleWhileRevalidate({cacheName:"static-font-assets",plugins:[new e.ExpirationPlugin({maxEntries:4,maxAgeSeconds:604800})]}),"GET"),e.registerRoute(/\.(?:jpg|jpeg|gif|png|svg|ico|webp)$/i,new e.StaleWhileRevalidate({cacheName:"static-image-assets",plugins:[new e.ExpirationPlugin({maxEntries:64,maxAgeSeconds:2592e3})]}),"GET"),e.registerRoute(/\/_next\/static.+\.js$/i,new e.CacheFirst({cacheName:"next-static-js-assets",plugins:[new e.ExpirationPlugin({maxEntries:64,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\/_next\/image\?url=.+$/i,new e.StaleWhileRevalidate({cacheName:"next-image",plugins:[new e.ExpirationPlugin({maxEntries:64,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:mp3|wav|ogg)$/i,new e.CacheFirst({cacheName:"static-audio-assets",plugins:[new e.RangeRequestsPlugin,new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:mp4|webm)$/i,new e.CacheFirst({cacheName:"static-video-assets",plugins:[new e.RangeRequestsPlugin,new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:js)$/i,new e.StaleWhileRevalidate({cacheName:"static-js-assets",plugins:[new e.ExpirationPlugin({maxEntries:48,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:css|less)$/i,new e.StaleWhileRevalidate({cacheName:"static-style-assets",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\/_next\/data\/.+\/.+\.json$/i,new e.StaleWhileRevalidate({cacheName:"next-data",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:json|xml|csv)$/i,new e.NetworkFirst({cacheName:"static-data-assets",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute((({sameOrigin:e,url:{pathname:s}})=>!(!e||s.startsWith("/api/auth/")||!s.startsWith("/api/"))),new e.NetworkFirst({cacheName:"apis",networkTimeoutSeconds:10,plugins:[new e.ExpirationPlugin({maxEntries:16,maxAgeSeconds:86400})]}),"GET"),e.registerRoute((({request:e,url:{pathname:s},sameOrigin:a})=>"1"===e.headers.get("RSC")&&"1"===e.headers.get("Next-Router-Prefetch")&&a&&!s.startsWith("/api/")),new e.NetworkFirst({cacheName:"pages-rsc-prefetch",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute((({request:e,url:{pathname:s},sameOrigin:a})=>"1"===e.headers.get("RSC")&&a&&!s.startsWith("/api/")),new e.NetworkFirst({cacheName:"pages-rsc",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute((({url:{pathname:e},sameOrigin:s})=>s&&!e.startsWith("/api/")),new e.NetworkFirst({cacheName:"pages",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute((({sameOrigin:e})=>!e),new e.NetworkFirst({cacheName:"cross-origin",networkTimeoutSeconds:10,plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:3600})]}),"GET")}));
//# sourceMappingURL=sw.js.map
