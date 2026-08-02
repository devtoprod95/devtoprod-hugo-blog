const cacheName = self.location.pathname
const pages = [

  "/devtoprod-hugo-blog/blog/20260802111809315/",
  "/devtoprod-hugo-blog/",
  "/devtoprod-hugo-blog/posts/",
  "/devtoprod-hugo-blog/blog/20260802111450230/",
  "/devtoprod-hugo-blog/blog/20260802113107582/",
  "/devtoprod-hugo-blog/blog/20260801230107536/",
  "/devtoprod-hugo-blog/blog/20260802053049465/",
  "/devtoprod-hugo-blog/blog/20260801190010150/",
  "/devtoprod-hugo-blog/blog/20260802005241052/",
  "/devtoprod-hugo-blog/blog/20260801215651055/",
  "/devtoprod-hugo-blog/blog/20260801220415131/",
  "/devtoprod-hugo-blog/blog/20260801221247247/",
  "/devtoprod-hugo-blog/blog/20260801154536064/",
  "/devtoprod-hugo-blog/blog/20260801071652723/",
  "/devtoprod-hugo-blog/blog/20260801184843089/",
  "/devtoprod-hugo-blog/blog/20260801133551705/",
  "/devtoprod-hugo-blog/blog/20260801150107539/",
  "/devtoprod-hugo-blog/blog/20260801112330529/",
  "/devtoprod-hugo-blog/blog/20260801090120718/",
  "/devtoprod-hugo-blog/blog/20260801111851483/",
  "/devtoprod-hugo-blog/blog/20260801113209641/",
  "/devtoprod-hugo-blog/blog/20260801053210700/",
  "/devtoprod-hugo-blog/blog/20260731230209761/",
  "/devtoprod-hugo-blog/blog/20260731214652456/",
  "/devtoprod-hugo-blog/blog/20260731225736710/",
  "/devtoprod-hugo-blog/blog/20260731225611679/",
  "/devtoprod-hugo-blog/blog/20260731184708132/",
  "/devtoprod-hugo-blog/blog/20260731190209579/",
  "/devtoprod-hugo-blog/blog/20260731153336536/",
  "/devtoprod-hugo-blog/blog/20260731113209490/",
  "/devtoprod-hugo-blog/blog/20260731072102597/",
  "/devtoprod-hugo-blog/blog/20260731111707859/",
  "/devtoprod-hugo-blog/blog/20260731060336361/",
  "/devtoprod-hugo-blog/blog/20260731010851587/",
  "/devtoprod-hugo-blog/blog/20260731013447793/",
  "/devtoprod-hugo-blog/blog/20260730223050128/",
  "/devtoprod-hugo-blog/blog/20260730224809353/",
  "/devtoprod-hugo-blog/blog/20260730180143772/",
  "/devtoprod-hugo-blog/blog/20260730190927525/",
  "/devtoprod-hugo-blog/blog/20260730191515714/",
  "/devtoprod-hugo-blog/blog/20260730143243571/",
  "/devtoprod-hugo-blog/blog/20260730104209782/",
  "/devtoprod-hugo-blog/blog/20260730152451676/",
  "/devtoprod-hugo-blog/blog/20260730152410649/",
  "/devtoprod-hugo-blog/blog/20260730105209309/",
  "/devtoprod-hugo-blog/blog/20260730030128343/",
  "/devtoprod-hugo-blog/blog/20260730012928049/",
  "/devtoprod-hugo-blog/blog/20260729225649097/",
  "/devtoprod-hugo-blog/blog/20260729225910140/",
  "/devtoprod-hugo-blog/blog/20260729183409297/",
  "/devtoprod-hugo-blog/blog/20260729192614538/",
  "/devtoprod-hugo-blog/blog/20260729193127654/",
  "/devtoprod-hugo-blog/blog/20260729152611635/",
  "/devtoprod-hugo-blog/blog/20260729111912723/",
  "/devtoprod-hugo-blog/blog/20260729105137109/",
  "/devtoprod-hugo-blog/blog/20260729112307909/",
  "/devtoprod-hugo-blog/blog/20260728232708840/",
  "/devtoprod-hugo-blog/blog/20260729060323461/",
  "/devtoprod-hugo-blog/blog/20260729010508970/",
  "/devtoprod-hugo-blog/blog/20260729010128940/",
  "/devtoprod-hugo-blog/blog/20260729011810069/",
  "/devtoprod-hugo-blog/blog/20260728225313581/",
  "/devtoprod-hugo-blog/blog/20260728225938625/",
  "/devtoprod-hugo-blog/blog/project-introduction/",
  "/devtoprod-hugo-blog/blog/20260728181249835/",
  "/devtoprod-hugo-blog/blog/20260728192754806/",
  "/devtoprod-hugo-blog/blog/20260728153109335/",
  "/devtoprod-hugo-blog/blog/20260728170337861/",
  "/devtoprod-hugo-blog/blog/20260728170855091/",
  "/devtoprod-hugo-blog/blog/20260728150307881/",
  "/devtoprod-hugo-blog/blog/20260728111948829/",
  "/devtoprod-hugo-blog/blog/20260728103334347/",
  "/devtoprod-hugo-blog/blog/20260728170610973/",
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