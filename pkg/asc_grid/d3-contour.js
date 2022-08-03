"use strict";
function _createForOfIteratorHelper(o, f) {
  var s = (typeof Symbol < "u" && o[Symbol.iterator]) || o["@@iterator"];
  if (!s) {
    if (
      Array.isArray(o) ||
      (s = _unsupportedIterableToArray(o)) ||
      (f && o && typeof o.length == "number")
    ) {
      s && (o = s);
      var L = 0,
        k = function () {};
      return {
        s: k,
        n: function () {
          return L >= o.length ? { done: !0 } : { done: !1, value: o[L++] };
        },
        e: function (z) {
          throw z;
        },
        f: k,
      };
    }
    throw new TypeError(`Invalid attempt to iterate non-iterable instance.
In order to be iterable, non-array objects must have a [Symbol.iterator]() method.`);
  }
  var U = !0,
    D = !1,
    H;
  return {
    s: function () {
      s = s.call(o);
    },
    n: function () {
      var z = s.next();
      return (U = z.done), z;
    },
    e: function (z) {
      (D = !0), (H = z);
    },
    f: function () {
      try {
        !U && s.return != null && s.return();
      } finally {
        if (D) throw H;
      }
    },
  };
}
function _unsupportedIterableToArray(o, f) {
  if (!!o) {
    if (typeof o == "string") return _arrayLikeToArray(o, f);
    var s = Object.prototype.toString.call(o).slice(8, -1);
    if (
      (s === "Object" && o.constructor && (s = o.constructor.name),
      s === "Map" || s === "Set")
    )
      return Array.from(o);
    if (s === "Arguments" || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(s))
      return _arrayLikeToArray(o, f);
  }
}
function _arrayLikeToArray(o, f) {
  (f == null || f > o.length) && (f = o.length);
  for (var s = 0, L = new Array(f); s < f; s++) L[s] = o[s];
  return L;
}
function _typeof(o) {
  return (
    (_typeof =
      typeof Symbol == "function" && typeof Symbol.iterator == "symbol"
        ? function (f) {
            return typeof f;
          }
        : function (f) {
            return f &&
              typeof Symbol == "function" &&
              f.constructor === Symbol &&
              f !== Symbol.prototype
              ? "symbol"
              : typeof f;
          }),
    _typeof(o)
  );
}
(function (o, f) {
  f(
    ((o = typeof globalThis < "u" ? globalThis : o || self).d3 = o.d3 || {}),
    o.d3
  );
})(void 0, function (o, f) {
  "use strict";
  var s = Array.prototype.slice;
  function L(e, h) {
    return e - h;
  }
  var k = function (h) {
    return function () {
      return h;
    };
  };
  function U(e, h) {
    for (var O, w = -1, d = h.length; ++w < d; ) if ((O = D(e, h[w]))) return O;
    return 0;
  }
  function D(e, h) {
    for (
      var O = h[0], w = h[1], d = -1, A = 0, g = e.length, v = g - 1;
      A < g;
      v = A++
    ) {
      var n = e[A],
        a = n[0],
        y = n[1],
        M = e[v],
        r = M[0],
        u = M[1];
      if (H(n, M, h)) return 0;
      y > w != u > w && O < ((r - a) * (w - y)) / (u - y) + a && (d = -d);
    }
    return d;
  }
  function H(e, h, O) {
    var w, d, A, g;
    return (
      (function (v, n, a) {
        return (n[0] - v[0]) * (a[1] - v[1]) == (a[0] - v[0]) * (n[1] - v[1]);
      })(e, h, O) &&
      ((d = e[(w = +(e[0] === h[0]))]),
      (A = O[w]),
      (g = h[w]),
      (d <= A && A <= g) || (g <= A && A <= d))
    );
  }
  function C() {}
  var z = [
    [],
    [
      [
        [1, 1.5],
        [0.5, 1],
      ],
    ],
    [
      [
        [1.5, 1],
        [1, 1.5],
      ],
    ],
    [
      [
        [1.5, 1],
        [0.5, 1],
      ],
    ],
    [
      [
        [1, 0.5],
        [1.5, 1],
      ],
    ],
    [
      [
        [1, 1.5],
        [0.5, 1],
      ],
      [
        [1, 0.5],
        [1.5, 1],
      ],
    ],
    [
      [
        [1, 0.5],
        [1, 1.5],
      ],
    ],
    [
      [
        [1, 0.5],
        [0.5, 1],
      ],
    ],
    [
      [
        [0.5, 1],
        [1, 0.5],
      ],
    ],
    [
      [
        [1, 1.5],
        [1, 0.5],
      ],
    ],
    [
      [
        [0.5, 1],
        [1, 0.5],
      ],
      [
        [1.5, 1],
        [1, 1.5],
      ],
    ],
    [
      [
        [1.5, 1],
        [1, 0.5],
      ],
    ],
    [
      [
        [0.5, 1],
        [1.5, 1],
      ],
    ],
    [
      [
        [1, 1.5],
        [1.5, 1],
      ],
    ],
    [
      [
        [0.5, 1],
        [1, 1.5],
      ],
    ],
    [],
  ];
  function _() {
    var e = 1,
      h = 1,
      O = f.thresholdSturges,
      w = v;
    function d(n) {
      var a = O(n);
      if (Array.isArray(a)) a = a.slice().sort(L);
      else {
        var y = f.extent(n),
          M = f.tickStep(y[0], y[1], a);
        a = f.ticks(Math.floor(y[0] / M) * M, Math.floor(y[1] / M - 1) * M, a);
      }
      return a.map(function (r) {
        return A(n, r);
      });
    }
    function A(n, a) {
      var y = [],
        M = [];
      return (
        (function (r, u, m) {
          var l,
            b,
            j,
            t,
            c,
            E,
            S = new Array(),
            p = new Array();
          for (l = b = -1, t = r[0] >= u, z[t << 1].forEach(I); ++l < e - 1; )
            (j = t), (t = r[l + 1] >= u), z[j | (t << 1)].forEach(I);
          for (z[t << 0].forEach(I); ++b < h - 1; ) {
            for (
              l = -1,
                t = r[b * e + e] >= u,
                c = r[b * e] >= u,
                z[(t << 1) | (c << 2)].forEach(I);
              ++l < e - 1;

            )
              (j = t),
                (t = r[b * e + e + l + 1] >= u),
                (E = c),
                (c = r[b * e + l + 1] >= u),
                z[j | (t << 1) | (c << 2) | (E << 3)].forEach(I);
            z[t | (c << 3)].forEach(I);
          }
          for (l = -1, c = r[b * e] >= u, z[c << 2].forEach(I); ++l < e - 1; )
            (E = c),
              (c = r[b * e + l + 1] >= u),
              z[(c << 2) | (E << 3)].forEach(I);
          function I(T) {
            var i,
              x,
              F = [T[0][0] + l, T[0][1] + b],
              q = [T[1][0] + l, T[1][1] + b],
              N = g(F),
              P = g(q);
            (i = p[N])
              ? (x = S[P])
                ? (delete p[i.end],
                  delete S[x.start],
                  i === x
                    ? (i.ring.push(q), m(i.ring))
                    : (S[i.start] = p[x.end] =
                        {
                          start: i.start,
                          end: x.end,
                          ring: i.ring.concat(x.ring),
                        }))
                : (delete p[i.end], i.ring.push(q), (p[(i.end = P)] = i))
              : (i = S[P])
              ? (x = p[N])
                ? (delete S[i.start],
                  delete p[x.end],
                  i === x
                    ? (i.ring.push(q), m(i.ring))
                    : (S[x.start] = p[i.end] =
                        {
                          start: x.start,
                          end: i.end,
                          ring: x.ring.concat(i.ring),
                        }))
                : (delete S[i.start], i.ring.unshift(F), (S[(i.start = N)] = i))
              : (S[N] = p[P] = { start: N, end: P, ring: [F, q] });
          }
          z[c << 3].forEach(I);
        })(n, a, function (r) {
          w(r, n, a),
            (function (u) {
              for (
                var m = 0,
                  l = u.length,
                  b = u[l - 1][1] * u[0][0] - u[l - 1][0] * u[0][1];
                ++m < l;

              )
                b += u[m - 1][1] * u[m][0] - u[m - 1][0] * u[m][1];
              return b;
            })(r) > 0
              ? y.push([r])
              : M.push(r);
        }),
        M.forEach(function (r) {
          for (var u, m = 0, l = y.length; m < l; ++m)
            if (U((u = y[m])[0], r) !== -1) return void u.push(r);
        }),
        { type: "MultiPolygon", value: a, coordinates: y }
      );
    }
    function g(n) {
      return 2 * n[0] + n[1] * (e + 1) * 4;
    }
    function v(n, a, y) {
      n.forEach(function (M) {
        var r,
          u = M[0],
          m = M[1],
          l = 0 | u,
          b = 0 | m,
          j = a[b * e + l];
        u > 0 &&
          u < e &&
          l === u &&
          ((r = a[b * e + l - 1]), (M[0] = u + (y - r) / (j - r) - 0.5)),
          m > 0 &&
            m < h &&
            b === m &&
            ((r = a[(b - 1) * e + l]), (M[1] = m + (y - r) / (j - r) - 0.5));
      });
    }
    return (
      (d.contour = A),
      (d.size = function (n) {
        if (!arguments.length) return [e, h];
        var a = Math.floor(n[0]),
          y = Math.floor(n[1]);
        if (!(a >= 0 && y >= 0)) throw new Error("invalid size");
        return (e = a), (h = y), d;
      }),
      (d.thresholds = function (n) {
        return arguments.length
          ? ((O =
              typeof n == "function"
                ? n
                : Array.isArray(n)
                ? k(s.call(n))
                : k(n)),
            d)
          : O;
      }),
      (d.smooth = function (n) {
        return arguments.length ? ((w = n ? v : C), d) : w === v;
      }),
      d
    );
  }

 
    (o.contours = _),
    Object.defineProperty(o, "__esModule", { value: !0 });
});
