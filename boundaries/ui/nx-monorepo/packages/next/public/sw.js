if(!self.define){let e,s={};const n=(n,a)=>(n=new URL(n+".js",a).href,s[n]||new Promise((s=>{if("document"in self){const e=document.createElement("script");e.src=n,e.onload=s,document.head.appendChild(e)}else e=n,importScripts(n),s()})).then((()=>{let e=s[n];if(!e)throw new Error(`Module ${n} didn’t register its module`);return e})));self.define=(a,t)=>{const i=e||("document"in self?document.currentScript.src:"")||location.href;if(s[i])return;let c={};const r=e=>n(e,i),d={module:{uri:i},exports:c,require:r};s[i]=Promise.all(a.map((e=>d[e]||r(e)))).then((e=>(t(...e),c)))}}define(["./workbox-9564d7f6"],(function(e){"use strict";importScripts(),self.skipWaiting(),e.clientsClaim(),e.precacheAndRoute([{url:"/next/_next/static/WWMD-Ty95fAWqv0lvaUuk/_buildManifest.js",revision:"548ae2728ac134bee787b4c22205dbf7"},{url:"/next/_next/static/WWMD-Ty95fAWqv0lvaUuk/_ssgManifest.js",revision:"b6652df95db52feb4daf4eca35380933"},{url:"/next/_next/static/chunks/384-52fa118dddf54ecd.js",revision:"52fa118dddf54ecd"},{url:"/next/_next/static/chunks/546-18f1a481e52a045e.js",revision:"18f1a481e52a045e"},{url:"/next/_next/static/chunks/563-a76d62f00ff9b751.js",revision:"a76d62f00ff9b751"},{url:"/next/_next/static/chunks/605-bbbf9c980df0ef2e.js",revision:"bbbf9c980df0ef2e"},{url:"/next/_next/static/chunks/684-5fadc2334fa455a1.js",revision:"5fadc2334fa455a1"},{url:"/next/_next/static/chunks/739-89cc35e9f7c05b8f.js",revision:"89cc35e9f7c05b8f"},{url:"/next/_next/static/chunks/785-9af97924ef827cf5.js",revision:"9af97924ef827cf5"},{url:"/next/_next/static/chunks/796-2d6df983f401eecf.js",revision:"2d6df983f401eecf"},{url:"/next/_next/static/chunks/dc2a7c46-85de9439db7c2777.js",revision:"85de9439db7c2777"},{url:"/next/_next/static/chunks/f69bbb46-5f16137fffb0f379.js",revision:"5f16137fffb0f379"},{url:"/next/_next/static/chunks/framework-fa45b7f8dd9b6f8e.js",revision:"fa45b7f8dd9b6f8e"},{url:"/next/_next/static/chunks/main-d26df60fb661bae6.js",revision:"d26df60fb661bae6"},{url:"/next/_next/static/chunks/pages/_app-b97c56d6a15b321c.js",revision:"b97c56d6a15b321c"},{url:"/next/_next/static/chunks/pages/_error-763b1ba722b25154.js",revision:"763b1ba722b25154"},{url:"/next/_next/static/chunks/pages/about-fc341b419db24c7f.js",revision:"fc341b419db24c7f"},{url:"/next/_next/static/chunks/pages/admin/groups-ea75f94202f7aefc.js",revision:"ea75f94202f7aefc"},{url:"/next/_next/static/chunks/pages/admin/links-8577f93dea3596c8.js",revision:"8577f93dea3596c8"},{url:"/next/_next/static/chunks/pages/admin/users-df80608f32063231.js",revision:"df80608f32063231"},{url:"/next/_next/static/chunks/pages/auth/forgot-815ba65f2dd12449.js",revision:"815ba65f2dd12449"},{url:"/next/_next/static/chunks/pages/auth/login-3c51a10e671b2f2d.js",revision:"3c51a10e671b2f2d"},{url:"/next/_next/static/chunks/pages/auth/registration-69bbb19478ffc8fe.js",revision:"69bbb19478ffc8fe"},{url:"/next/_next/static/chunks/pages/auth/verification-d3f1bd137040d188.js",revision:"d3f1bd137040d188"},{url:"/next/_next/static/chunks/pages/contact-9b002342e967e0da.js",revision:"9b002342e967e0da"},{url:"/next/_next/static/chunks/pages/faq-7efe809723e8db02.js",revision:"7efe809723e8db02"},{url:"/next/_next/static/chunks/pages/index-522950d72d40934b.js",revision:"522950d72d40934b"},{url:"/next/_next/static/chunks/pages/pricing-8b63a2dec0302ba2.js",revision:"8b63a2dec0302ba2"},{url:"/next/_next/static/chunks/pages/privacy-3549624762807fe0.js",revision:"3549624762807fe0"},{url:"/next/_next/static/chunks/pages/user/addUrl-a39259de28e64d27.js",revision:"a39259de28e64d27"},{url:"/next/_next/static/chunks/pages/user/audit-e1714d307a536990.js",revision:"e1714d307a536990"},{url:"/next/_next/static/chunks/pages/user/billing-0d7cc06cf9230836.js",revision:"0d7cc06cf9230836"},{url:"/next/_next/static/chunks/pages/user/dashboard-22e4902cc69b309f.js",revision:"22e4902cc69b309f"},{url:"/next/_next/static/chunks/pages/user/integrations-d7043bed19a6036c.js",revision:"d7043bed19a6036c"},{url:"/next/_next/static/chunks/pages/user/links-83eca65e7b1588ca.js",revision:"83eca65e7b1588ca"},{url:"/next/_next/static/chunks/pages/user/profile-7f4e66a1b594c6e2.js",revision:"7f4e66a1b594c6e2"},{url:"/next/_next/static/chunks/pages/user/reports-db7c879561dec66d.js",revision:"db7c879561dec66d"},{url:"/next/_next/static/chunks/pages/user/security-076ed0dfda836dca.js",revision:"076ed0dfda836dca"},{url:"/next/_next/static/chunks/polyfills-c67a75d1b6f99dc8.js",revision:"837c0df77fd5009c9e46d446188ecfd0"},{url:"/next/_next/static/chunks/webpack-8c891280b3e5141a.js",revision:"8c891280b3e5141a"},{url:"/next/_next/static/css/16fa344f79acb6a9.css",revision:"16fa344f79acb6a9"},{url:"/next/assets/images/undraw_back_in_the_day_knsh.svg",revision:"aebc6c499a138c3e107e65a208aec647"},{url:"/next/assets/images/undraw_co_workers_re_1i6i.svg",revision:"cb908c2f6d43c3d5bced6e0804dac2e9"},{url:"/next/assets/images/undraw_designer_re_5v95.svg",revision:"435c0b4cb909d0ceb63048a4e7ebc9f5"},{url:"/next/assets/images/undraw_welcome_cats_thqn.svg",revision:"ed0c3358facded075949f5e0ab20a080"},{url:"/next/assets/styles.css",revision:"edc2f8f4dfae29085e43adcb26dbe64b"},{url:"/next/favicon.ico",revision:"c30c7d42707a47a3f4591831641e50dc"},{url:"/next/firebase-messaging-sw.js",revision:"47db0543b0c9d21608ee0cda826ce944"},{url:"/next/manifest.json",revision:"44354e9e77eae0d74431df55175e5566"},{url:"/next/sitemap-0.xml",revision:"2ef8f99fa4fba2551f9effced9738793"}],{ignoreURLParametersMatching:[/^utm_/,/^fbclid$/]}),e.cleanupOutdatedCaches(),e.registerRoute("/next",new e.NetworkFirst({cacheName:"start-url",plugins:[{cacheWillUpdate:async({response:e})=>e&&"opaqueredirect"===e.type?new Response(e.body,{status:200,statusText:"OK",headers:e.headers}):e}]}),"GET"),e.registerRoute(/^https:\/\/fonts\.(?:gstatic)\.com\/.*/i,new e.CacheFirst({cacheName:"google-fonts-webfonts",plugins:[new e.ExpirationPlugin({maxEntries:4,maxAgeSeconds:31536e3})]}),"GET"),e.registerRoute(/^https:\/\/fonts\.(?:googleapis)\.com\/.*/i,new e.StaleWhileRevalidate({cacheName:"google-fonts-stylesheets",plugins:[new e.ExpirationPlugin({maxEntries:4,maxAgeSeconds:604800})]}),"GET"),e.registerRoute(/\.(?:eot|otf|ttc|ttf|woff|woff2|font.css)$/i,new e.StaleWhileRevalidate({cacheName:"static-font-assets",plugins:[new e.ExpirationPlugin({maxEntries:4,maxAgeSeconds:604800})]}),"GET"),e.registerRoute(/\.(?:jpg|jpeg|gif|png|svg|ico|webp)$/i,new e.StaleWhileRevalidate({cacheName:"static-image-assets",plugins:[new e.ExpirationPlugin({maxEntries:64,maxAgeSeconds:2592e3})]}),"GET"),e.registerRoute(/\/_next\/static.+\.js$/i,new e.CacheFirst({cacheName:"next-static-js-assets",plugins:[new e.ExpirationPlugin({maxEntries:64,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\/_next\/image\?url=.+$/i,new e.StaleWhileRevalidate({cacheName:"next-image",plugins:[new e.ExpirationPlugin({maxEntries:64,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:mp3|wav|ogg)$/i,new e.CacheFirst({cacheName:"static-audio-assets",plugins:[new e.RangeRequestsPlugin,new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:mp4|webm)$/i,new e.CacheFirst({cacheName:"static-video-assets",plugins:[new e.RangeRequestsPlugin,new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:js)$/i,new e.StaleWhileRevalidate({cacheName:"static-js-assets",plugins:[new e.ExpirationPlugin({maxEntries:48,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:css|less)$/i,new e.StaleWhileRevalidate({cacheName:"static-style-assets",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\/_next\/data\/.+\/.+\.json$/i,new e.StaleWhileRevalidate({cacheName:"next-data",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:json|xml|csv)$/i,new e.NetworkFirst({cacheName:"static-data-assets",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute((({sameOrigin:e,url:{pathname:s}})=>!(!e||s.startsWith("/api/auth/callback")||!s.startsWith("/api/"))),new e.NetworkFirst({cacheName:"apis",networkTimeoutSeconds:10,plugins:[new e.ExpirationPlugin({maxEntries:16,maxAgeSeconds:86400})]}),"GET"),e.registerRoute((({request:e,url:{pathname:s},sameOrigin:n})=>"1"===e.headers.get("RSC")&&"1"===e.headers.get("Next-Router-Prefetch")&&n&&!s.startsWith("/api/")),new e.NetworkFirst({cacheName:"pages-rsc-prefetch",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute((({request:e,url:{pathname:s},sameOrigin:n})=>"1"===e.headers.get("RSC")&&n&&!s.startsWith("/api/")),new e.NetworkFirst({cacheName:"pages-rsc",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute((({url:{pathname:e},sameOrigin:s})=>s&&!e.startsWith("/api/")),new e.NetworkFirst({cacheName:"pages",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute((({sameOrigin:e})=>!e),new e.NetworkFirst({cacheName:"cross-origin",networkTimeoutSeconds:10,plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:3600})]}),"GET")}));
//# sourceMappingURL=sw.js.map
