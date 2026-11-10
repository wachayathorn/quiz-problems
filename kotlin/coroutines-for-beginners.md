# Kotlin Coroutines for Beginners

Simple English guide. Short sentences. Many examples.

You need only basic Kotlin first: functions, classes, and `val` / `var`.

---

## 1. What is a coroutine?

A **coroutine** is a small task that can **pause** and **continue later**.

It can wait for something (like a network call) without freezing the whole program.

Think like this:

- A normal function runs from start to end. It does not pause in a smart way.
- A coroutine can say: “I wait here. You can do other work. Call me again later.”

### Simple picture

```text
Normal blocking wait:
  Thread: [work][WAIT........][work]
           (this thread is busy while waiting)

Coroutine wait:
  Thread: [work][free for other tasks][work again]
           (the wait does not hold the thread the whole time)
```

---

## 2. Why do we use coroutines?

### Problem: blocking code

```kotlin
fun loadUser(): String {
    Thread.sleep(2000) // bad in apps: the thread sits and does nothing
    return "Ann"
}
```

`Thread.sleep` makes the thread wait for 2 seconds. That thread cannot help other work.

### Better idea with coroutines

```kotlin
suspend fun loadUser(): String {
    delay(2000) // wait without blocking the thread like sleep
    return "Ann"
}
```

`delay` is a **suspend** wait. The coroutine pauses. The thread can do other work.

### When coroutines help

- Call APIs
- Read files or databases
- Run many waits at the same time
- Keep UI / server threads free

---

## 3. The word `suspend`

`suspend` means: this function **can pause**.

```kotlin
suspend fun hello(): String {
    delay(500)
    return "hello"
}
```

### Rules (easy version)

1. A `suspend` function can call other `suspend` functions.
2. A normal function **cannot** call a `suspend` function directly.
3. You need a **coroutine** to start `suspend` code.

So this is wrong:

```kotlin
fun main() {
    hello() // error: suspend function from normal function
}
```

This is OK:

```kotlin
fun main() = runBlocking {
    println(hello())
}
```

---

## 4. How do we start a coroutine?

You need a **builder**. Builders create and start coroutines.

### 4.1 `runBlocking`

`runBlocking` starts a coroutine and **blocks** the current thread until the work finishes.

Use it in:

- `main`
- simple tests

Do **not** use it inside real app services for every call.

```kotlin
import kotlinx.coroutines.*

fun main() = runBlocking {
    println("start")
    delay(1000)
    println("end")
}
```

Output:

```text
start
(wait 1 second)
end
```

### 4.2 `launch`

`launch` starts a coroutine that does work. It does **not** return a useful result value.

It returns a `Job` (a handle for the task).

```kotlin
fun main() = runBlocking {
    launch {
        delay(500)
        println("from launch")
    }
    println("from main")
}
```

Possible output order:

```text
from main
from launch
```

Why? `launch` starts work and continues. The next line can run first.

### 4.3 `async` and `await`

`async` starts a coroutine that **returns a value**.

It returns `Deferred<T>`.  
Call `await()` to get the value.

```kotlin
fun main() = runBlocking {
    val deferred: Deferred<Int> = async {
        delay(300)
        42
    }
    println("waiting...")
    val value = deferred.await()
    println(value)
}
```

Output:

```text
waiting...
42
```

### `launch` vs `async`

| Builder | Returns | Use when |
|---|---|---|
| `launch` | `Job` | You only want to run work |
| `async` | `Deferred<T>` | You want a result later |

---

## 5. Run many tasks at the same time

This is a common reason to use coroutines.

```kotlin
fun main() = runBlocking {
    val time = measureTimeMillis {
        val a = async {
            delay(1000)
            "A"
        }
        val b = async {
            delay(1000)
            "B"
        }
        println(a.await() + b.await())
    }
    println("time = $time ms")
}
```

Both waits are about 1 second. Together they finish near **1 second**, not 2 seconds, because they run **in parallel**.

If you write this instead (one after one):

```kotlin
fun main() = runBlocking {
    val a = doA() // waits 1s
    val b = doB() // waits 1s more
}

suspend fun doA(): String {
    delay(1000)
    return "A"
}

suspend fun doB(): String {
    delay(1000)
    return "B"
}
```

Total wait is about **2 seconds**.

---

## 6. Dispatchers: where does the code run?

A **dispatcher** chooses the thread pool for your coroutine.

### Common dispatchers

| Dispatcher | Simple meaning |
|---|---|
| `Dispatchers.Main` | UI thread (Android) |
| `Dispatchers.IO` | Wait work: files, network, DB |
| `Dispatchers.Default` | Heavy CPU work |
| `Dispatchers.Unconfined` | Advanced. Skip for now. |

### Example

```kotlin
fun main() = runBlocking {
    launch(Dispatchers.Default) {
        println("CPU work on ${Thread.currentThread().name}")
    }
    launch(Dispatchers.IO) {
        println("IO work on ${Thread.currentThread().name}")
    }
}
```

You do not need to memorize thread names. Remember the idea: **pick the right place for the work**.

---

## 7. `withContext`: move work, then come back

`withContext` runs a block on another dispatcher, waits for it, and returns the result.

```kotlin
suspend fun readFileName(): String = withContext(Dispatchers.IO) {
    // pretend this is slow file work
    delay(200)
    "notes.txt"
}

fun main() = runBlocking {
    val name = readFileName()
    println(name)
}
```

### When to use `withContext`

- One piece of work
- You want the result now
- You want to move only that piece to `IO` or `Default`

### `withContext` vs `async`

```kotlin
// one task: withContext is clear
val data = withContext(Dispatchers.IO) { load() }

// many tasks at the same time: async is better
coroutineScope {
    val a = async(Dispatchers.IO) { loadA() }
    val b = async(Dispatchers.IO) { loadB() }
    a.await() to b.await()
}
```

---

## 8. `coroutineScope` and structured concurrency

**Structured concurrency** means: child tasks belong to a parent scope.  
When the parent ends, children should end too. Errors are easier to see.

### `coroutineScope`

```kotlin
suspend fun loadScreen(): String = coroutineScope {
    val user = async { loadUser() }
    val posts = async { loadPosts() }
    "${user.await()} / ${posts.await()}"
}

suspend fun loadUser(): String {
    delay(200)
    return "Ann"
}

suspend fun loadPosts(): String {
    delay(200)
    return "3 posts"
}
```

What `coroutineScope` does:

1. Starts a safe scope for children.
2. Waits until children finish.
3. If one child fails, other children are cancelled.
4. Then the error goes up.

### Do not use `GlobalScope` as a beginner

```kotlin
GlobalScope.launch { ... } // hard to control. avoid this first
```

Prefer `coroutineScope`, `runBlocking` (for demos), or scopes from Android / Spring.

---

## 9. Cancellation and timeout

### Cancel a job

```kotlin
fun main() = runBlocking {
    val job = launch {
        repeat(1000) { i ->
            println("i = $i")
            delay(200) // delay checks cancel
        }
    }
    delay(500)
    println("cancel now")
    job.cancel()
    job.join() // wait until the job really stops
    println("done")
}
```

Important idea: cancellation works well at **suspend points** like `delay` and `await`.

### Timeout

```kotlin
fun main() = runBlocking {
    try {
        withTimeout(300) {
            delay(1000)
            println("too slow")
        }
    } catch (e: TimeoutCancellationException) {
        println("timeout!")
    }
}
```

Or get `null` instead of an exception:

```kotlin
fun main() = runBlocking {
    val result = withTimeoutOrNull(300) {
        delay(1000)
        "ok"
    }
    println(result) // null
}
```

---

## 10. Flow basics (many values over time)

A normal `suspend` function returns **one** value.

A **Flow** can send **many** values, one by one, over time.

### Make a flow

```kotlin
fun numbers(): Flow<Int> = flow {
    emit(1)
    delay(200)
    emit(2)
    delay(200)
    emit(3)
}
```

- `flow { }` creates the stream.
- `emit(x)` sends one value.

### Collect values

```kotlin
fun main() = runBlocking {
    numbers().collect { value ->
        println(value)
    }
}
```

Output:

```text
1
(wait)
2
(wait)
3
```

`collect` is a `suspend` function. It waits for each value.

### Map and filter

```kotlin
fun main() = runBlocking {
    numbers()
        .filter { it % 2 == 1 } // keep odd numbers
        .map { "n=$it" }       // change each value
        .collect { println(it) }
}
```

Output:

```text
n=1
n=3
```

### Flow vs list

| | List | Flow |
|---|---|---|
| When values exist | All at once in memory | One by one over time |
| Good for | Small known data | Streams, events, live updates |

Cold flow note (simple): the work inside `flow { }` starts when you call `collect`.

---

## 11. Mini cheat sheet

```kotlin
// pause function
suspend fun f() { delay(100) }

// start and block current thread (demo / test)
runBlocking { ... }

// start work, no result value
launch { ... }

// start work with result
val d = async { 1 }
val x = d.await()

// move one block to IO
withContext(Dispatchers.IO) { ... }

// parent scope for children
coroutineScope {
    val a = async { ... }
    val b = async { ... }
    a.await(); b.await()
}

// timeout
withTimeout(1_000) { ... }

// many values over time
flow { emit(1) }.collect { println(it) }
```

---

## 12. Common beginner mistakes

1. Call `suspend` from a normal function with no builder.  
2. Use `Thread.sleep` inside coroutines. Prefer `delay`.  
3. Use only one `async` and think it is “parallel”. Parallel needs two or more tasks together.  
4. Use `GlobalScope` everywhere. Prefer a clear scope.  
5. Forget `await()` after `async` when you need the value.

---

## 13. Tiny practice

Try this in a Kotlin project with `kotlinx-coroutines-core`:

```kotlin
import kotlinx.coroutines.*
import kotlin.system.measureTimeMillis

suspend fun fetchA(): String {
    delay(800)
    return "A"
}

suspend fun fetchB(): String {
    delay(800)
    return "B"
}

fun main() = runBlocking {
    val sequential = measureTimeMillis {
        println(fetchA() + fetchB())
    }
    println("sequential ms = $sequential")

    val parallel = measureTimeMillis {
        val result = coroutineScope {
            val a = async { fetchA() }
            val b = async { fetchB() }
            a.await() + b.await()
        }
        println(result)
    }
    println("parallel ms = $parallel")
}
```

You should see:

- sequential time near **1600+** ms
- parallel time near **800+** ms

---

## 14. What to learn next

After this file:

1. Use coroutines in Android with `viewModelScope` / `lifecycleScope`
2. Use `suspend` in Spring controllers and services
3. Learn more Flow operators (`flatMapConcat`, `combine`, …)
4. Learn `supervisorScope` when one child fail should not cancel siblings

---

## End

Remember three words:

1. **`suspend`** — this function can pause  
2. **builder** — `runBlocking` / `launch` / `async` starts work  
3. **scope** — parent owns children (`coroutineScope`)

Start small. Run the examples. Change the delays and watch the order of prints.
