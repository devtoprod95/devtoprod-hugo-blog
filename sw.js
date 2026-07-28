const cacheName = self.location.pathname
const pages = [

  "/devtoprod-hugo-blog/",
  "/devtoprod-hugo-blog/posts/",
  "/devtoprod-hugo-blog/blog/20260727172313261/",
  "/devtoprod-hugo-blog/blog/20260728153109335/",
  "/devtoprod-hugo-blog/blog/20260728163151459/",
  "/devtoprod-hugo-blog/blog/20260728163710653/",
  "/devtoprod-hugo-blog/blog/2026-07-28-153051-naver-ranking/",
  "/devtoprod-hugo-blog/blog/2026-07-28-153040-naver-relationship/",
  "/devtoprod-hugo-blog/blog/2026-07-28-153031-naver-music/",
  "/devtoprod-hugo-blog/blog/2026-07-28-153021-naver-movie/",
  "/devtoprod-hugo-blog/blog/2026-07-28-153012-naver-drama/",
  "/devtoprod-hugo-blog/blog/project-introduction/",
  "/devtoprod-hugo-blog/blog/etc-sample/",
  "/devtoprod-hugo-blog/blog/book-sample/",
  "/devtoprod-hugo-blog/blog/dev-sample/",
  "/devtoprod-hugo-blog/categories/",
  "/devtoprod-hugo-blog/list/",
  "/devtoprod-hugo-blog/list/newest/",
  "/devtoprod-hugo-blog/list/oldest/",
  "/devtoprod-hugo-blog/tags/",
  "/devtoprod-hugo-blog/search/",
  "/devtoprod-hugo-blog/main.min.f9c55459aa1fb5a9ce5b6ae28b563e5f91180c6a8549b7a74acdcf81e3977f28.css",
  
];

self.addEventListener("install", function (event) {
  self.skipWaiting();

  caches.open(cacheName).then((cache) => {
    return cache.addAll(pages);
  });
});

self.addEventListener("fetch", (event) => {
  const request = event.request;
  if (request.method !== "GET") {
    return;
  }

  /**
   * @param {Response} response
   * @returns {Promise<Response>}
   */
  function saveToCache(response) {
    if (cacheable(response)) {
      return caches
        .open(cacheName)
        .then((cache) => cache.put(request, response.clone()))
        .then(() => response);
    } else {
      return response;
    }
  }

  /**
   * @param {Error} error
   */
  function serveFromCache(error) {
    return caches.open(cacheName).then((cache) => cache.match(request.url));
  }

  /**
   * @param {Response} response
   * @returns {Boolean}
   */
  function cacheable(response) {
    return response.type === "basic" && response.ok && !response.headers.has("Content-Disposition")
  }

  event.respondWith(fetch(request).then(saveToCache).catch(serveFromCache));
});