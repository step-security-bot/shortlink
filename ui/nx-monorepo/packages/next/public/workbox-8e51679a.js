define(['exports'], function (t) {
  'use strict'
  try {
    self['workbox:core:7.0.0'] && _()
  } catch (t) {}
  const e = (t, ...e) => {
    let s = t
    return e.length > 0 && (s += ` :: ${JSON.stringify(e)}`), s
  }
  class s extends Error {
    constructor(t, s) {
      super(e(t, s)), (this.name = t), (this.details = s)
    }
  }
  try {
    self['workbox:routing:7.0.0'] && _()
  } catch (t) {}
  const n = (t) => (t && 'object' == typeof t ? t : { handle: t })
  class r {
    constructor(t, e, s = 'GET') {
      ;(this.handler = n(e)), (this.match = t), (this.method = s)
    }
    setCatchHandler(t) {
      this.catchHandler = n(t)
    }
  }
  class i extends r {
    constructor(t, e, s) {
      super(
        ({ url: e }) => {
          const s = t.exec(e.href)
          if (s && (e.origin === location.origin || 0 === s.index))
            return s.slice(1)
        },
        e,
        s,
      )
    }
  }
  class a {
    constructor() {
      ;(this.t = new Map()), (this.i = new Map())
    }
    get routes() {
      return this.t
    }
    addFetchListener() {
      self.addEventListener('fetch', (t) => {
        const { request: e } = t,
          s = this.handleRequest({ request: e, event: t })
        s && t.respondWith(s)
      })
    }
    addCacheListener() {
      self.addEventListener('message', (t) => {
        if (t.data && 'CACHE_URLS' === t.data.type) {
          const { payload: e } = t.data,
            s = Promise.all(
              e.urlsToCache.map((e) => {
                'string' == typeof e && (e = [e])
                const s = new Request(...e)
                return this.handleRequest({ request: s, event: t })
              }),
            )
          t.waitUntil(s),
            t.ports && t.ports[0] && s.then(() => t.ports[0].postMessage(!0))
        }
      })
    }
    handleRequest({ request: t, event: e }) {
      const s = new URL(t.url, location.href)
      if (!s.protocol.startsWith('http')) return
      const n = s.origin === location.origin,
        { params: r, route: i } = this.findMatchingRoute({
          event: e,
          request: t,
          sameOrigin: n,
          url: s,
        })
      let a = i && i.handler
      const o = t.method
      if ((!a && this.i.has(o) && (a = this.i.get(o)), !a)) return
      let c
      try {
        c = a.handle({ url: s, request: t, event: e, params: r })
      } catch (t) {
        c = Promise.reject(t)
      }
      const h = i && i.catchHandler
      return (
        c instanceof Promise &&
          (this.o || h) &&
          (c = c.catch(async (n) => {
            if (h)
              try {
                return await h.handle({
                  url: s,
                  request: t,
                  event: e,
                  params: r,
                })
              } catch (t) {
                t instanceof Error && (n = t)
              }
            if (this.o) return this.o.handle({ url: s, request: t, event: e })
            throw n
          })),
        c
      )
    }
    findMatchingRoute({ url: t, sameOrigin: e, request: s, event: n }) {
      const r = this.t.get(s.method) || []
      for (const i of r) {
        let r
        const a = i.match({ url: t, sameOrigin: e, request: s, event: n })
        if (a)
          return (
            (r = a),
            ((Array.isArray(r) && 0 === r.length) ||
              (a.constructor === Object && 0 === Object.keys(a).length) ||
              'boolean' == typeof a) &&
              (r = void 0),
            { route: i, params: r }
          )
      }
      return {}
    }
    setDefaultHandler(t, e = 'GET') {
      this.i.set(e, n(t))
    }
    setCatchHandler(t) {
      this.o = n(t)
    }
    registerRoute(t) {
      this.t.has(t.method) || this.t.set(t.method, []),
        this.t.get(t.method).push(t)
    }
    unregisterRoute(t) {
      if (!this.t.has(t.method))
        throw new s('unregister-route-but-not-found-with-method', {
          method: t.method,
        })
      const e = this.t.get(t.method).indexOf(t)
      if (!(e > -1)) throw new s('unregister-route-route-not-registered')
      this.t.get(t.method).splice(e, 1)
    }
  }
  let o
  const c = () => (
    o || ((o = new a()), o.addFetchListener(), o.addCacheListener()), o
  )
  function h(t, e, n) {
    let a
    if ('string' == typeof t) {
      const s = new URL(t, location.href)
      a = new r(({ url: t }) => t.href === s.href, e, n)
    } else if (t instanceof RegExp) a = new i(t, e, n)
    else if ('function' == typeof t) a = new r(t, e, n)
    else {
      if (!(t instanceof r))
        throw new s('unsupported-route-type', {
          moduleName: 'workbox-routing',
          funcName: 'registerRoute',
          paramName: 'capture',
        })
      a = t
    }
    return c().registerRoute(a), a
  }
  try {
    self['workbox:strategies:7.0.0'] && _()
  } catch (t) {}
  const u = {
      cacheWillUpdate: async ({ response: t }) =>
        200 === t.status || 0 === t.status ? t : null,
    },
    l = {
      googleAnalytics: 'googleAnalytics',
      precache: 'precache-v2',
      prefix: 'workbox',
      runtime: 'runtime',
      suffix: 'undefined' != typeof registration ? registration.scope : '',
    },
    f = (t) =>
      [l.prefix, t, l.suffix].filter((t) => t && t.length > 0).join('-'),
    w = (t) => t || f(l.precache),
    d = (t) => t || f(l.runtime)
  function p(t, e) {
    const s = new URL(t)
    for (const t of e) s.searchParams.delete(t)
    return s.href
  }
  class y {
    constructor() {
      this.promise = new Promise((t, e) => {
        ;(this.resolve = t), (this.reject = e)
      })
    }
  }
  const g = new Set()
  function m(t) {
    return 'string' == typeof t ? new Request(t) : t
  }
  class R {
    constructor(t, e) {
      ;(this.h = {}),
        Object.assign(this, e),
        (this.event = e.event),
        (this.u = t),
        (this.l = new y()),
        (this.p = []),
        (this.g = [...t.plugins]),
        (this.m = new Map())
      for (const t of this.g) this.m.set(t, {})
      this.event.waitUntil(this.l.promise)
    }
    async fetch(t) {
      const { event: e } = this
      let n = m(t)
      if (
        'navigate' === n.mode &&
        e instanceof FetchEvent &&
        e.preloadResponse
      ) {
        const t = await e.preloadResponse
        if (t) return t
      }
      const r = this.hasCallback('fetchDidFail') ? n.clone() : null
      try {
        for (const t of this.iterateCallbacks('requestWillFetch'))
          n = await t({ request: n.clone(), event: e })
      } catch (t) {
        if (t instanceof Error)
          throw new s('plugin-error-request-will-fetch', {
            thrownErrorMessage: t.message,
          })
      }
      const i = n.clone()
      try {
        let t
        t = await fetch(n, 'navigate' === n.mode ? void 0 : this.u.fetchOptions)
        for (const s of this.iterateCallbacks('fetchDidSucceed'))
          t = await s({ event: e, request: i, response: t })
        return t
      } catch (t) {
        throw (
          (r &&
            (await this.runCallbacks('fetchDidFail', {
              error: t,
              event: e,
              originalRequest: r.clone(),
              request: i.clone(),
            })),
          t)
        )
      }
    }
    async fetchAndCachePut(t) {
      const e = await this.fetch(t),
        s = e.clone()
      return this.waitUntil(this.cachePut(t, s)), e
    }
    async cacheMatch(t) {
      const e = m(t)
      let s
      const { cacheName: n, matchOptions: r } = this.u,
        i = await this.getCacheKey(e, 'read'),
        a = Object.assign(Object.assign({}, r), { cacheName: n })
      s = await caches.match(i, a)
      for (const t of this.iterateCallbacks('cachedResponseWillBeUsed'))
        s =
          (await t({
            cacheName: n,
            matchOptions: r,
            cachedResponse: s,
            request: i,
            event: this.event,
          })) || void 0
      return s
    }
    async cachePut(t, e) {
      const n = m(t)
      var r
      await ((r = 0), new Promise((t) => setTimeout(t, r)))
      const i = await this.getCacheKey(n, 'write')
      if (!e)
        throw new s('cache-put-with-no-response', {
          url:
            ((a = i.url),
            new URL(String(a), location.href).href.replace(
              new RegExp(`^${location.origin}`),
              '',
            )),
        })
      var a
      const o = await this.R(e)
      if (!o) return !1
      const { cacheName: c, matchOptions: h } = this.u,
        u = await self.caches.open(c),
        l = this.hasCallback('cacheDidUpdate'),
        f = l
          ? await (async function (t, e, s, n) {
              const r = p(e.url, s)
              if (e.url === r) return t.match(e, n)
              const i = Object.assign(Object.assign({}, n), {
                  ignoreSearch: !0,
                }),
                a = await t.keys(e, i)
              for (const e of a) if (r === p(e.url, s)) return t.match(e, n)
            })(u, i.clone(), ['__WB_REVISION__'], h)
          : null
      try {
        await u.put(i, l ? o.clone() : o)
      } catch (t) {
        if (t instanceof Error)
          throw (
            ('QuotaExceededError' === t.name &&
              (await (async function () {
                for (const t of g) await t()
              })()),
            t)
          )
      }
      for (const t of this.iterateCallbacks('cacheDidUpdate'))
        await t({
          cacheName: c,
          oldResponse: f,
          newResponse: o.clone(),
          request: i,
          event: this.event,
        })
      return !0
    }
    async getCacheKey(t, e) {
      const s = `${t.url} | ${e}`
      if (!this.h[s]) {
        let n = t
        for (const t of this.iterateCallbacks('cacheKeyWillBeUsed'))
          n = m(
            await t({
              mode: e,
              request: n,
              event: this.event,
              params: this.params,
            }),
          )
        this.h[s] = n
      }
      return this.h[s]
    }
    hasCallback(t) {
      for (const e of this.u.plugins) if (t in e) return !0
      return !1
    }
    async runCallbacks(t, e) {
      for (const s of this.iterateCallbacks(t)) await s(e)
    }
    *iterateCallbacks(t) {
      for (const e of this.u.plugins)
        if ('function' == typeof e[t]) {
          const s = this.m.get(e),
            n = (n) => {
              const r = Object.assign(Object.assign({}, n), { state: s })
              return e[t](r)
            }
          yield n
        }
    }
    waitUntil(t) {
      return this.p.push(t), t
    }
    async doneWaiting() {
      let t
      for (; (t = this.p.shift()); ) await t
    }
    destroy() {
      this.l.resolve(null)
    }
    async R(t) {
      let e = t,
        s = !1
      for (const t of this.iterateCallbacks('cacheWillUpdate'))
        if (
          ((e =
            (await t({
              request: this.request,
              response: e,
              event: this.event,
            })) || void 0),
          (s = !0),
          !e)
        )
          break
      return s || (e && 200 !== e.status && (e = void 0)), e
    }
  }
  class v {
    constructor(t = {}) {
      ;(this.cacheName = d(t.cacheName)),
        (this.plugins = t.plugins || []),
        (this.fetchOptions = t.fetchOptions),
        (this.matchOptions = t.matchOptions)
    }
    handle(t) {
      const [e] = this.handleAll(t)
      return e
    }
    handleAll(t) {
      t instanceof FetchEvent && (t = { event: t, request: t.request })
      const e = t.event,
        s = 'string' == typeof t.request ? new Request(t.request) : t.request,
        n = 'params' in t ? t.params : void 0,
        r = new R(this, { event: e, request: s, params: n }),
        i = this.v(r, s, e)
      return [i, this.q(i, r, s, e)]
    }
    async v(t, e, n) {
      let r
      await t.runCallbacks('handlerWillStart', { event: n, request: e })
      try {
        if (((r = await this.D(e, t)), !r || 'error' === r.type))
          throw new s('no-response', { url: e.url })
      } catch (s) {
        if (s instanceof Error)
          for (const i of t.iterateCallbacks('handlerDidError'))
            if (((r = await i({ error: s, event: n, request: e })), r)) break
        if (!r) throw s
      }
      for (const s of t.iterateCallbacks('handlerWillRespond'))
        r = await s({ event: n, request: e, response: r })
      return r
    }
    async q(t, e, s, n) {
      let r, i
      try {
        r = await t
      } catch (i) {}
      try {
        await e.runCallbacks('handlerDidRespond', {
          event: n,
          request: s,
          response: r,
        }),
          await e.doneWaiting()
      } catch (t) {
        t instanceof Error && (i = t)
      }
      if (
        (await e.runCallbacks('handlerDidComplete', {
          event: n,
          request: s,
          response: r,
          error: i,
        }),
        e.destroy(),
        i)
      )
        throw i
    }
  }
  function b(t) {
    t.then(() => {})
  }
  function q() {
    return (
      (q = Object.assign
        ? Object.assign.bind()
        : function (t) {
            for (var e = 1; e < arguments.length; e++) {
              var s = arguments[e]
              for (var n in s)
                Object.prototype.hasOwnProperty.call(s, n) && (t[n] = s[n])
            }
            return t
          }),
      q.apply(this, arguments)
    )
  }
  const D = (t, e) => e.some((e) => t instanceof e)
  let U, x
  const L = new WeakMap(),
    I = new WeakMap(),
    C = new WeakMap(),
    E = new WeakMap(),
    N = new WeakMap()
  let O = {
    get(t, e, s) {
      if (t instanceof IDBTransaction) {
        if ('done' === e) return I.get(t)
        if ('objectStoreNames' === e) return t.objectStoreNames || C.get(t)
        if ('store' === e)
          return s.objectStoreNames[1]
            ? void 0
            : s.objectStore(s.objectStoreNames[0])
      }
      return B(t[e])
    },
    set: (t, e, s) => ((t[e] = s), !0),
    has: (t, e) =>
      (t instanceof IDBTransaction && ('done' === e || 'store' === e)) ||
      e in t,
  }
  function T(t) {
    return t !== IDBDatabase.prototype.transaction ||
      'objectStoreNames' in IDBTransaction.prototype
      ? (
          x ||
          (x = [
            IDBCursor.prototype.advance,
            IDBCursor.prototype.continue,
            IDBCursor.prototype.continuePrimaryKey,
          ])
        ).includes(t)
        ? function (...e) {
            return t.apply(P(this), e), B(L.get(this))
          }
        : function (...e) {
            return B(t.apply(P(this), e))
          }
      : function (e, ...s) {
          const n = t.call(P(this), e, ...s)
          return C.set(n, e.sort ? e.sort() : [e]), B(n)
        }
  }
  function k(t) {
    return 'function' == typeof t
      ? T(t)
      : (t instanceof IDBTransaction &&
          (function (t) {
            if (I.has(t)) return
            const e = new Promise((e, s) => {
              const n = () => {
                  t.removeEventListener('complete', r),
                    t.removeEventListener('error', i),
                    t.removeEventListener('abort', i)
                },
                r = () => {
                  e(), n()
                },
                i = () => {
                  s(t.error || new DOMException('AbortError', 'AbortError')),
                    n()
                }
              t.addEventListener('complete', r),
                t.addEventListener('error', i),
                t.addEventListener('abort', i)
            })
            I.set(t, e)
          })(t),
        D(
          t,
          U ||
            (U = [
              IDBDatabase,
              IDBObjectStore,
              IDBIndex,
              IDBCursor,
              IDBTransaction,
            ]),
        )
          ? new Proxy(t, O)
          : t)
  }
  function B(t) {
    if (t instanceof IDBRequest)
      return (function (t) {
        const e = new Promise((e, s) => {
          const n = () => {
              t.removeEventListener('success', r),
                t.removeEventListener('error', i)
            },
            r = () => {
              e(B(t.result)), n()
            },
            i = () => {
              s(t.error), n()
            }
          t.addEventListener('success', r), t.addEventListener('error', i)
        })
        return (
          e
            .then((e) => {
              e instanceof IDBCursor && L.set(e, t)
            })
            .catch(() => {}),
          N.set(e, t),
          e
        )
      })(t)
    if (E.has(t)) return E.get(t)
    const e = k(t)
    return e !== t && (E.set(t, e), N.set(e, t)), e
  }
  const P = (t) => N.get(t)
  const M = ['get', 'getKey', 'getAll', 'getAllKeys', 'count'],
    W = ['put', 'add', 'delete', 'clear'],
    j = new Map()
  function S(t, e) {
    if (!(t instanceof IDBDatabase) || e in t || 'string' != typeof e) return
    if (j.get(e)) return j.get(e)
    const s = e.replace(/FromIndex$/, ''),
      n = e !== s,
      r = W.includes(s)
    if (
      !(s in (n ? IDBIndex : IDBObjectStore).prototype) ||
      (!r && !M.includes(s))
    )
      return
    const i = async function (t, ...e) {
      const i = this.transaction(t, r ? 'readwrite' : 'readonly')
      let a = i.store
      return (
        n && (a = a.index(e.shift())),
        (await Promise.all([a[s](...e), r && i.done]))[0]
      )
    }
    return j.set(e, i), i
  }
  O = ((t) =>
    q({}, t, {
      get: (e, s, n) => S(e, s) || t.get(e, s, n),
      has: (e, s) => !!S(e, s) || t.has(e, s),
    }))(O)
  try {
    self['workbox:expiration:7.0.0'] && _()
  } catch (t) {}
  const K = 'cache-entries',
    A = (t) => {
      const e = new URL(t, location.href)
      return (e.hash = ''), e.href
    }
  class F {
    constructor(t) {
      ;(this.U = null), (this._ = t)
    }
    L(t) {
      const e = t.createObjectStore(K, { keyPath: 'id' })
      e.createIndex('cacheName', 'cacheName', { unique: !1 }),
        e.createIndex('timestamp', 'timestamp', { unique: !1 })
    }
    I(t) {
      this.L(t),
        this._ &&
          (function (t, { blocked: e } = {}) {
            const s = indexedDB.deleteDatabase(t)
            e && s.addEventListener('blocked', (t) => e(t.oldVersion, t)),
              B(s).then(() => {})
          })(this._)
    }
    async setTimestamp(t, e) {
      const s = {
          url: (t = A(t)),
          timestamp: e,
          cacheName: this._,
          id: this.C(t),
        },
        n = (await this.getDb()).transaction(K, 'readwrite', {
          durability: 'relaxed',
        })
      await n.store.put(s), await n.done
    }
    async getTimestamp(t) {
      const e = await this.getDb(),
        s = await e.get(K, this.C(t))
      return null == s ? void 0 : s.timestamp
    }
    async expireEntries(t, e) {
      const s = await this.getDb()
      let n = await s
        .transaction(K)
        .store.index('timestamp')
        .openCursor(null, 'prev')
      const r = []
      let i = 0
      for (; n; ) {
        const s = n.value
        s.cacheName === this._ &&
          ((t && s.timestamp < t) || (e && i >= e) ? r.push(n.value) : i++),
          (n = await n.continue())
      }
      const a = []
      for (const t of r) await s.delete(K, t.id), a.push(t.url)
      return a
    }
    C(t) {
      return this._ + '|' + A(t)
    }
    async getDb() {
      return (
        this.U ||
          (this.U = await (function (
            t,
            e,
            { blocked: s, upgrade: n, blocking: r, terminated: i } = {},
          ) {
            const a = indexedDB.open(t, e),
              o = B(a)
            return (
              n &&
                a.addEventListener('upgradeneeded', (t) => {
                  n(
                    B(a.result),
                    t.oldVersion,
                    t.newVersion,
                    B(a.transaction),
                    t,
                  )
                }),
              s &&
                a.addEventListener('blocked', (t) =>
                  s(t.oldVersion, t.newVersion, t),
                ),
              o
                .then((t) => {
                  i && t.addEventListener('close', () => i()),
                    r &&
                      t.addEventListener('versionchange', (t) =>
                        r(t.oldVersion, t.newVersion, t),
                      )
                })
                .catch(() => {}),
              o
            )
          })('workbox-expiration', 1, { upgrade: this.I.bind(this) })),
        this.U
      )
    }
  }
  class H {
    constructor(t, e = {}) {
      ;(this.N = !1),
        (this.O = !1),
        (this.T = e.maxEntries),
        (this.k = e.maxAgeSeconds),
        (this.B = e.matchOptions),
        (this._ = t),
        (this.P = new F(t))
    }
    async expireEntries() {
      if (this.N) return void (this.O = !0)
      this.N = !0
      const t = this.k ? Date.now() - 1e3 * this.k : 0,
        e = await this.P.expireEntries(t, this.T),
        s = await self.caches.open(this._)
      for (const t of e) await s.delete(t, this.B)
      ;(this.N = !1), this.O && ((this.O = !1), b(this.expireEntries()))
    }
    async updateTimestamp(t) {
      await this.P.setTimestamp(t, Date.now())
    }
    async isURLExpired(t) {
      if (this.k) {
        const e = await this.P.getTimestamp(t),
          s = Date.now() - 1e3 * this.k
        return void 0 === e || e < s
      }
      return !1
    }
    async delete() {
      ;(this.O = !1), await this.P.expireEntries(1 / 0)
    }
  }
  try {
    self['workbox:range-requests:7.0.0'] && _()
  } catch (t) {}
  async function $(t, e) {
    try {
      if (206 === e.status) return e
      const n = t.headers.get('range')
      if (!n) throw new s('no-range-header')
      const r = (function (t) {
          const e = t.trim().toLowerCase()
          if (!e.startsWith('bytes='))
            throw new s('unit-must-be-bytes', { normalizedRangeHeader: e })
          if (e.includes(','))
            throw new s('single-range-only', { normalizedRangeHeader: e })
          const n = /(\d*)-(\d*)/.exec(e)
          if (!n || (!n[1] && !n[2]))
            throw new s('invalid-range-values', { normalizedRangeHeader: e })
          return {
            start: '' === n[1] ? void 0 : Number(n[1]),
            end: '' === n[2] ? void 0 : Number(n[2]),
          }
        })(n),
        i = await e.blob(),
        a = (function (t, e, n) {
          const r = t.size
          if ((n && n > r) || (e && e < 0))
            throw new s('range-not-satisfiable', { size: r, end: n, start: e })
          let i, a
          return (
            void 0 !== e && void 0 !== n
              ? ((i = e), (a = n + 1))
              : void 0 !== e && void 0 === n
              ? ((i = e), (a = r))
              : void 0 !== n && void 0 === e && ((i = r - n), (a = r)),
            { start: i, end: a }
          )
        })(i, r.start, r.end),
        o = i.slice(a.start, a.end),
        c = o.size,
        h = new Response(o, {
          status: 206,
          statusText: 'Partial Content',
          headers: e.headers,
        })
      return (
        h.headers.set('Content-Length', String(c)),
        h.headers.set(
          'Content-Range',
          `bytes ${a.start}-${a.end - 1}/${i.size}`,
        ),
        h
      )
    } catch (t) {
      return new Response('', {
        status: 416,
        statusText: 'Range Not Satisfiable',
      })
    }
  }
  function z(t, e) {
    const s = e()
    return t.waitUntil(s), s
  }
  try {
    self['workbox:precaching:7.0.0'] && _()
  } catch (t) {}
  function G(t) {
    if (!t) throw new s('add-to-cache-list-unexpected-type', { entry: t })
    if ('string' == typeof t) {
      const e = new URL(t, location.href)
      return { cacheKey: e.href, url: e.href }
    }
    const { revision: e, url: n } = t
    if (!n) throw new s('add-to-cache-list-unexpected-type', { entry: t })
    if (!e) {
      const t = new URL(n, location.href)
      return { cacheKey: t.href, url: t.href }
    }
    const r = new URL(n, location.href),
      i = new URL(n, location.href)
    return (
      r.searchParams.set('__WB_REVISION__', e),
      { cacheKey: r.href, url: i.href }
    )
  }
  class V {
    constructor() {
      ;(this.updatedURLs = []),
        (this.notUpdatedURLs = []),
        (this.handlerWillStart = async ({ request: t, state: e }) => {
          e && (e.originalRequest = t)
        }),
        (this.cachedResponseWillBeUsed = async ({
          event: t,
          state: e,
          cachedResponse: s,
        }) => {
          if (
            'install' === t.type &&
            e &&
            e.originalRequest &&
            e.originalRequest instanceof Request
          ) {
            const t = e.originalRequest.url
            s ? this.notUpdatedURLs.push(t) : this.updatedURLs.push(t)
          }
          return s
        })
    }
  }
  class J {
    constructor({ precacheController: t }) {
      ;(this.cacheKeyWillBeUsed = async ({ request: t, params: e }) => {
        const s =
          (null == e ? void 0 : e.cacheKey) || this.M.getCacheKeyForURL(t.url)
        return s ? new Request(s, { headers: t.headers }) : t
      }),
        (this.M = t)
    }
  }
  let Q, X
  async function Y(t, e) {
    let n = null
    if (t.url) {
      n = new URL(t.url).origin
    }
    if (n !== self.location.origin)
      throw new s('cross-origin-copy-response', { origin: n })
    const r = t.clone(),
      i = {
        headers: new Headers(r.headers),
        status: r.status,
        statusText: r.statusText,
      },
      a = e ? e(i) : i,
      o = (function () {
        if (void 0 === Q) {
          const t = new Response('')
          if ('body' in t)
            try {
              new Response(t.body), (Q = !0)
            } catch (t) {
              Q = !1
            }
          Q = !1
        }
        return Q
      })()
        ? r.body
        : await r.blob()
    return new Response(o, a)
  }
  class Z extends v {
    constructor(t = {}) {
      ;(t.cacheName = w(t.cacheName)),
        super(t),
        (this.W = !1 !== t.fallbackToNetwork),
        this.plugins.push(Z.copyRedirectedCacheableResponsesPlugin)
    }
    async D(t, e) {
      const s = await e.cacheMatch(t)
      return (
        s ||
        (e.event && 'install' === e.event.type
          ? await this.j(t, e)
          : await this.S(t, e))
      )
    }
    async S(t, e) {
      let n
      const r = e.params || {}
      if (!this.W)
        throw new s('missing-precache-entry', {
          cacheName: this.cacheName,
          url: t.url,
        })
      {
        const s = r.integrity,
          i = t.integrity,
          a = !i || i === s
        ;(n = await e.fetch(
          new Request(t, { integrity: 'no-cors' !== t.mode ? i || s : void 0 }),
        )),
          s &&
            a &&
            'no-cors' !== t.mode &&
            (this.K(), await e.cachePut(t, n.clone()))
      }
      return n
    }
    async j(t, e) {
      this.K()
      const n = await e.fetch(t)
      if (!(await e.cachePut(t, n.clone())))
        throw new s('bad-precaching-response', { url: t.url, status: n.status })
      return n
    }
    K() {
      let t = null,
        e = 0
      for (const [s, n] of this.plugins.entries())
        n !== Z.copyRedirectedCacheableResponsesPlugin &&
          (n === Z.defaultPrecacheCacheabilityPlugin && (t = s),
          n.cacheWillUpdate && e++)
      0 === e
        ? this.plugins.push(Z.defaultPrecacheCacheabilityPlugin)
        : e > 1 && null !== t && this.plugins.splice(t, 1)
    }
  }
  ;(Z.defaultPrecacheCacheabilityPlugin = {
    cacheWillUpdate: async ({ response: t }) =>
      !t || t.status >= 400 ? null : t,
  }),
    (Z.copyRedirectedCacheableResponsesPlugin = {
      cacheWillUpdate: async ({ response: t }) =>
        t.redirected ? await Y(t) : t,
    })
  class tt {
    constructor({
      cacheName: t,
      plugins: e = [],
      fallbackToNetwork: s = !0,
    } = {}) {
      ;(this.A = new Map()),
        (this.F = new Map()),
        (this.H = new Map()),
        (this.u = new Z({
          cacheName: w(t),
          plugins: [...e, new J({ precacheController: this })],
          fallbackToNetwork: s,
        })),
        (this.install = this.install.bind(this)),
        (this.activate = this.activate.bind(this))
    }
    get strategy() {
      return this.u
    }
    precache(t) {
      this.addToCacheList(t),
        this.$ ||
          (self.addEventListener('install', this.install),
          self.addEventListener('activate', this.activate),
          (this.$ = !0))
    }
    addToCacheList(t) {
      const e = []
      for (const n of t) {
        'string' == typeof n
          ? e.push(n)
          : n && void 0 === n.revision && e.push(n.url)
        const { cacheKey: t, url: r } = G(n),
          i = 'string' != typeof n && n.revision ? 'reload' : 'default'
        if (this.A.has(r) && this.A.get(r) !== t)
          throw new s('add-to-cache-list-conflicting-entries', {
            firstEntry: this.A.get(r),
            secondEntry: t,
          })
        if ('string' != typeof n && n.integrity) {
          if (this.H.has(t) && this.H.get(t) !== n.integrity)
            throw new s('add-to-cache-list-conflicting-integrities', { url: r })
          this.H.set(t, n.integrity)
        }
        if ((this.A.set(r, t), this.F.set(r, i), e.length > 0)) {
          const t = `Workbox is precaching URLs without revision info: ${e.join(
            ', ',
          )}\nThis is generally NOT safe. Learn more at https://bit.ly/wb-precache`
          console.warn(t)
        }
      }
    }
    install(t) {
      return z(t, async () => {
        const e = new V()
        this.strategy.plugins.push(e)
        for (const [e, s] of this.A) {
          const n = this.H.get(s),
            r = this.F.get(e),
            i = new Request(e, {
              integrity: n,
              cache: r,
              credentials: 'same-origin',
            })
          await Promise.all(
            this.strategy.handleAll({
              params: { cacheKey: s },
              request: i,
              event: t,
            }),
          )
        }
        const { updatedURLs: s, notUpdatedURLs: n } = e
        return { updatedURLs: s, notUpdatedURLs: n }
      })
    }
    activate(t) {
      return z(t, async () => {
        const t = await self.caches.open(this.strategy.cacheName),
          e = await t.keys(),
          s = new Set(this.A.values()),
          n = []
        for (const r of e) s.has(r.url) || (await t.delete(r), n.push(r.url))
        return { deletedURLs: n }
      })
    }
    getURLsToCacheKeys() {
      return this.A
    }
    getCachedURLs() {
      return [...this.A.keys()]
    }
    getCacheKeyForURL(t) {
      const e = new URL(t, location.href)
      return this.A.get(e.href)
    }
    getIntegrityForCacheKey(t) {
      return this.H.get(t)
    }
    async matchPrecache(t) {
      const e = t instanceof Request ? t.url : t,
        s = this.getCacheKeyForURL(e)
      if (s) {
        return (await self.caches.open(this.strategy.cacheName)).match(s)
      }
    }
    createHandlerBoundToURL(t) {
      const e = this.getCacheKeyForURL(t)
      if (!e) throw new s('non-precached-url', { url: t })
      return (s) => (
        (s.request = new Request(t)),
        (s.params = Object.assign({ cacheKey: e }, s.params)),
        this.strategy.handle(s)
      )
    }
  }
  const et = () => (X || (X = new tt()), X)
  class st extends r {
    constructor(t, e) {
      super(({ request: s }) => {
        const n = t.getURLsToCacheKeys()
        for (const r of (function* (
          t,
          {
            ignoreURLParametersMatching: e = [/^utm_/, /^fbclid$/],
            directoryIndex: s = 'index.html',
            cleanURLs: n = !0,
            urlManipulation: r,
          } = {},
        ) {
          const i = new URL(t, location.href)
          ;(i.hash = ''), yield i.href
          const a = (function (t, e = []) {
            for (const s of [...t.searchParams.keys()])
              e.some((t) => t.test(s)) && t.searchParams.delete(s)
            return t
          })(i, e)
          if ((yield a.href, s && a.pathname.endsWith('/'))) {
            const t = new URL(a.href)
            ;(t.pathname += s), yield t.href
          }
          if (n) {
            const t = new URL(a.href)
            ;(t.pathname += '.html'), yield t.href
          }
          if (r) {
            const t = r({ url: i })
            for (const e of t) yield e.href
          }
        })(s.url, e)) {
          const e = n.get(r)
          if (e) {
            return { cacheKey: e, integrity: t.getIntegrityForCacheKey(e) }
          }
        }
      }, t.strategy)
    }
  }
  ;(t.CacheFirst = class extends v {
    async D(t, e) {
      let n,
        r = await e.cacheMatch(t)
      if (!r)
        try {
          r = await e.fetchAndCachePut(t)
        } catch (t) {
          t instanceof Error && (n = t)
        }
      if (!r) throw new s('no-response', { url: t.url, error: n })
      return r
    }
  }),
    (t.ExpirationPlugin = class {
      constructor(t = {}) {
        ;(this.cachedResponseWillBeUsed = async ({
          event: t,
          request: e,
          cacheName: s,
          cachedResponse: n,
        }) => {
          if (!n) return null
          const r = this.G(n),
            i = this.V(s)
          b(i.expireEntries())
          const a = i.updateTimestamp(e.url)
          if (t)
            try {
              t.waitUntil(a)
            } catch (t) {}
          return r ? n : null
        }),
          (this.cacheDidUpdate = async ({ cacheName: t, request: e }) => {
            const s = this.V(t)
            await s.updateTimestamp(e.url), await s.expireEntries()
          }),
          (this.J = t),
          (this.k = t.maxAgeSeconds),
          (this.X = new Map()),
          t.purgeOnQuotaError &&
            (function (t) {
              g.add(t)
            })(() => this.deleteCacheAndMetadata())
      }
      V(t) {
        if (t === d()) throw new s('expire-custom-caches-only')
        let e = this.X.get(t)
        return e || ((e = new H(t, this.J)), this.X.set(t, e)), e
      }
      G(t) {
        if (!this.k) return !0
        const e = this.Y(t)
        if (null === e) return !0
        return e >= Date.now() - 1e3 * this.k
      }
      Y(t) {
        if (!t.headers.has('date')) return null
        const e = t.headers.get('date'),
          s = new Date(e).getTime()
        return isNaN(s) ? null : s
      }
      async deleteCacheAndMetadata() {
        for (const [t, e] of this.X)
          await self.caches.delete(t), await e.delete()
        this.X = new Map()
      }
    }),
    (t.NetworkFirst = class extends v {
      constructor(t = {}) {
        super(t),
          this.plugins.some((t) => 'cacheWillUpdate' in t) ||
            this.plugins.unshift(u),
          (this.Z = t.networkTimeoutSeconds || 0)
      }
      async D(t, e) {
        const n = [],
          r = []
        let i
        if (this.Z) {
          const { id: s, promise: a } = this.tt({
            request: t,
            logs: n,
            handler: e,
          })
          ;(i = s), r.push(a)
        }
        const a = this.et({ timeoutId: i, request: t, logs: n, handler: e })
        r.push(a)
        const o = await e.waitUntil(
          (async () => (await e.waitUntil(Promise.race(r))) || (await a))(),
        )
        if (!o) throw new s('no-response', { url: t.url })
        return o
      }
      tt({ request: t, logs: e, handler: s }) {
        let n
        return {
          promise: new Promise((e) => {
            n = setTimeout(async () => {
              e(await s.cacheMatch(t))
            }, 1e3 * this.Z)
          }),
          id: n,
        }
      }
      async et({ timeoutId: t, request: e, logs: s, handler: n }) {
        let r, i
        try {
          i = await n.fetchAndCachePut(e)
        } catch (t) {
          t instanceof Error && (r = t)
        }
        return t && clearTimeout(t), (!r && i) || (i = await n.cacheMatch(e)), i
      }
    }),
    (t.RangeRequestsPlugin = class {
      constructor() {
        this.cachedResponseWillBeUsed = async ({
          request: t,
          cachedResponse: e,
        }) => (e && t.headers.has('range') ? await $(t, e) : e)
      }
    }),
    (t.StaleWhileRevalidate = class extends v {
      constructor(t = {}) {
        super(t),
          this.plugins.some((t) => 'cacheWillUpdate' in t) ||
            this.plugins.unshift(u)
      }
      async D(t, e) {
        const n = e.fetchAndCachePut(t).catch(() => {})
        e.waitUntil(n)
        let r,
          i = await e.cacheMatch(t)
        if (i);
        else
          try {
            i = await n
          } catch (t) {
            t instanceof Error && (r = t)
          }
        if (!i) throw new s('no-response', { url: t.url, error: r })
        return i
      }
    }),
    (t.cleanupOutdatedCaches = function () {
      self.addEventListener('activate', (t) => {
        const e = w()
        t.waitUntil(
          (async (t, e = '-precache-') => {
            const s = (await self.caches.keys()).filter(
              (s) =>
                s.includes(e) && s.includes(self.registration.scope) && s !== t,
            )
            return await Promise.all(s.map((t) => self.caches.delete(t))), s
          })(e).then((t) => {}),
        )
      })
    }),
    (t.clientsClaim = function () {
      self.addEventListener('activate', () => self.clients.claim())
    }),
    (t.precacheAndRoute = function (t, e) {
      !(function (t) {
        et().precache(t)
      })(t),
        (function (t) {
          const e = et()
          h(new st(e, t))
        })(e)
    }),
    (t.registerRoute = h)
})
//# sourceMappingURL=workbox-8e51679a.js.map
