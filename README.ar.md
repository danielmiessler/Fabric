<div align="center">
    <a href="https://go.warp.dev/fabric" target="_blank">
        <sup>شكر خاص لـ:</sup>
        <br>
        <img alt="رعاية Warp" width="400" src="https://raw.githubusercontent.com/warpdotdev/brand-assets/refs/heads/main/Github/Sponsor/Warp-Github-LG-02.png">
        <br>
        <h>Warp، مبني للبرمجة مع وكلاء الذكاء الاصطناعي المتعددين</b>
        <br>
        <sup>متاح لأنظمة macOS، Linux و Windows</sup>
    </a>
</div>

<br>

<div align="center">

<img src="./docs/images/fabric-logo-gif.gif" alt="شعار fabric" width="400" height="400"/>

# `fabric`

[![Static Badge](https://img.shields.io/badge/mission-human_flourishing_via_AI_augmentation-purple)](https://github.com/danielmiessler/fabric)
<br />
[![GitHub top language](https://img.shields.io/github/languages/top/danielmiessler/fabric)](https://github.com/danielmiessler/fabric)
[![GitHub last commit](https://img.shields.io/github/last-commit/danielmiessler/fabric)](https://github.com/danielmiessler/fabric/commits/main)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/danielmiessler/fabric)

<div align="center">
<h4><code>fabric</code> هو إطار عمل مفتوح المصدر لتعزيز قدرات البشر باستخدام الذكاء الاصطناعي.</h4>
</div>

<p align="center">
  <strong>العربية</strong> ·
  <a href="README.md">الإنجليزية</a> ·
  <a href="README.zh.md">الصينية</a>
</p>

![لقطة شاشة لـ fabric](./docs/images/fabric-summarize.png)

</div>

[التحديثات](#التحديثات) •
[ما ولماذا](#ما-ولماذا) •
[الفلسفة](#الفلسفة) •
[التثبيت](#التثبيت) •
[الاستخدام](#الاستخدام) •
[واجهة برمجة التطبيقات REST](#واجهة-برمجة-التطبيقات-rest) •
[أمثلة](#أمثلة) •
[مجرد استخدام الأنماط](#مجرد-استخدام-الأنماط) •
[الأنماط المخصصة](#الأنماط-المخصصة) •
[التطبيقات المساعدة](#التطبيقات-المساعدة) •
[معلومات إضافية](#معلومات-إضافية)

</div>

## ماهو ولماذا

منذ بداية الذكاء الاصطناعي الحديث في أواخر عام 2022، رأينا عددًا **_استثنائيًا_** من تطبيقات الذكاء الاصطناعي لإنجاز المهام. هناك الآلاف من المواقع، روبوتات الدردشة، التطبيقات المحمولة، والواجهات الأخرى لاستخدام جميع أنواع الذكاء الاصطناعي المختلفة.

كل هذا مثير حقًا وقوي، لكن _ليس من السهل دمج هذه الوظائف في حياتنا._

<div class="align center">
<h4>بمعنى آخر، لا يعاني الذكاء الاصطناعي من مشكلة في القدرات—بل يعاني من مشكلة <em>تكامل</em>.</h4>
</div>

**تم إنشاء Fabric لمعالجة هذه المشكلة من خلال إنشاء وتنظيم الوحدات الأساسية للذكاء الاصطناعي—التروجات نفسها!**

ينظم Fabric التروجات حسب المهمة الواقعية، مما يسمح للأشخاص بإنشاء وجمع وتنظيم أهم حلول الذكاء الاصطناعي الخاصة بهم في مكان واحد للاستخدام في أدواتهم المفضلة. وإذا كنت تركز على سطر الأوامر، يمكنك استخدام Fabric نفسه كواجهة!

## التحديثات

للتعمق في Fabric وداخله، اقرأ الوثائق في [مجلد docs](https://github.com/danielmiessler/Fabric/tree/main/docs). هناك أيضًا [DeepWiki](https://deepwiki.com/danielmiessler/Fabric) المفيد للغاية والمحدث بانتظام لـ Fabric.

<details>
<summary>انقر لعرض التحديثات الأخيرة</summary>

أعزائي المستخدمين،

لقد كنا نقوم بالعديد من الأشياء المثيرة هنا في Fabric، وأردت أن أعطي ملخصًا سريعًا هنا لإعطائكم إحساسًا بسرعة تطورنا!

أدناه **الميزات والقدرات الجديدة** التي أضفناها (الأحدث أولاً):

### الميزات الرئيسية الأخيرة

- [v1.4.447](https://github.com/danielmiessler/fabric/releases/tag/v1.4.447) (16 أبريل 2026) — **Claude Opus 4.7**: تحديث SDK الخاص بـ Anthropic إلى الإصدار v1.37.0 وإضافة [Claude Opus 4.7](https://www.anthropic.com/news/claude-opus-4-7) الجديد إلى النماذج المتاحة، بما في ذلك دعم نافذة سياق 1 مليون رمز.
- [v1.4.437](https://github.com/danielmiessler/fabric/releases/tag/v1.4.437) (16 مارس 2026) — **إضافة OpenAI Codex**: يدعم Fabric الآن استخدام OpenAI Codex (مع اشتراك OpenAI الخاص بك) كخلفية!
- [v1.4.417](https://github.com/danielmiessler/fabric/releases/tag/v1.4.417) (21 فبراير 2026) — **إضافة Azure AI Gateway**: إضافة إضافة Azure AI Gateway التي تدعم خلفيات متعددة (AWS Bedrock، Azure OpenAI، Google Vertex AI) من خلال بوابة Azure APIM الموحدة مع مصادقة مفتاح اشتراك مشترك.
- [v1.4.416](https://github.com/danielmiessler/fabric/releases/tag/v1.4.416) (21 فبراير 2026) — **مصادقة Azure Entra ID**: إضافة إضافة مصادقة Azure Entra ID مع أدوات Azure المشتركة، ودعم Entra ID/MSAL، واستخراج منطق Azure المشترك إلى حزمة قابلة لإعادة الاستخدام `azurecommon`.
- [v1.4.380](https://github.com/danielmiessler/fabric/releases/tag/v1.4.380) (15 يناير 2026) — **تكامل Microsoft 365 Copilot**: إضافة دعم Microsoft 365 Copilot المؤسسي، مما يتيح للمستخدمين المؤسسيين الاستفادة من الذكاء الاصطناعي المستند إلى بيانات مؤسستهم في Microsoft 365 (البريد الإلكتروني، المستندات، الاجتماعات).
- [v1.4.378](https://github.com/danielmiessler/fabric/releases/tag/v1.4.378) (14 يناير 2026) — **دعم Digital Ocean GenAI**: إضافة دعم Digital Ocean GenAI، مع [دليل حول كيفية استخدامه](./docs/DigitalOcean-Agents-Setup.md).
- [v1.4.356](https://github.com/danielmiessler/fabric/releases/tag/v1.4.356) (22 ديسمبر 2025) — **التدويل الكامل**: دعم i18n كامل لتروجات الإعداد عبر جميع اللغات العشر مع معالجة ذكية لمتغيرات البيئة—مما يجعل Fabric متاحًا حقًا في جميع أنحاء العالم مع الحفاظ على اتساق التكوين.
- [v1.4.350](https://github.com/danielmiessler/fabric/releases/tag/v1.4.350) (18 ديسمبر 2025) — **وثائق API التفاعلية**: إضافة واجهة Swagger/OpenAPI في `/swagger/index.html` مع وثائق REST API شاملة، وأدلة مطور محسنة، وتحسين اكتشاف نقاط النهاية لتكامل أسهل.
- [v1.4.338](https://github.com/danielmiessler/fabric/releases/tag/v1.4.338) (4 ديسمبر 2025) — إضافة دعم بائع Abacus لنماذج Chat-LLM (راجع [RouteLLM APIs](https://abacus.ai/app/route-llm-apis)).
- [v1.4.337](https://github.com/danielmiessler/fabric/releases/tag/v1.4.337) (4 ديسمبر 2025) — إضافة دعم بائع "Z AI". راجع صفحة [نظرة عامة على Z AI](https://docs.z.ai/guides/overview/overview) لمزيد من التفاصيل.
- [v1.4.334](https://github.com/danielmiessler/fabric/releases/tag/v1.4.334) (26 نوفمبر 2025) — **Claude Opus 4.5**: تحديث SDK الخاص بـ Anthropic إلى أحدث إصدار وإضافة [Claude Opus 4.5](https://www.anthropic.com/news/claude-opus-4-5) الجديد إلى النماذج المتاحة.
- [v1.4.331](https://github.com/danielmiessler/fabric/releases/tag/v1.4.331) (23 نوفمبر 2025) — **دعم نماذج GitHub**: إضافة دعم استخدام نماذج GitHub.
- [v1.4.322](https://github.com/danielmiessler/fabric/releases/tag/v1.4.322) (5 نوفمبر 2025) — **خرائط المفاهيم التفاعلية HTML و Claude Sonnet 4.5**: إضافة نمط `create_conceptmap` للتمثيل البصري للمعرفة باستخدام Vis.js، تقديم فئة WELLNESS مع أنماط التحليل النفسي، وترقية إلى Claude Sonnet 4.5
- [v1.4.317](https://github.com/danielmiessler/fabric/releases/tag/v1.4.317) (21 سبتمبر 2025) — **متغيرات اللغة البرتغالية**: إضافة تطبيع BCP 47 للمنطقة مع دعم البرتغالية البرازيلية (pt-BR) والبرتغالية الأوروبية (pt-PT) مع سلاسل استرجاع ذكية
- [v1.4.314](https://github.com/danielmiessler/fabric/releases/tag/v1.4.314) (17 سبتمبر 2025) — **هجرة Azure OpenAI**: الهجرة إلى `openai-go/azure` SDK الرسمي مع تحسين المصادقة ودعم إصدار API الافتراضي

تمثل هذه الميزات التزامنا بجعل Fabric أقوى وأكثر مرونة إطار عمل تعزيز الذكاء الاصطناعي المتاح!

</details>

## مقاطع الفيديو التقديمية

ضع في اعتبارك أن العديد من هذه المقاطع تم تسجيلها عندما كان Fabric يعتمد على Python، لذا تذكر استخدام [تعليمات التثبيت الحالية](#التثبيت) أدناه.

- [Network Chuck](https://www.youtube.com/watch?v=UbDyjIIGaxQ)
- [David Bombal](https://www.youtube.com/watch?v=vF-MQmVxnCs)
- [المقدمة الخاصة بي للأداة](https://www.youtube.com/watch?v=wPEyyigh10g)
- [المزيد من مقاطع الفيديو على YouTube لـ Fabric](https://www.youtube.com/results?search_query=fabric+ai)

## الفلسفة

> الذكاء الاصطناعي ليس شيئًا؛ إنه _مكبر_ لشيء. وهذا الشيء هو **الإبداع البشري**.

نعتقد أن الغرض من التكنولوجيا هو مساعدة البشر على الازدهار، لذلك عندما نتحدث عن الذكاء الاصطناعي نبدأ بالمشاكل **البشرية** التي نريد حلها.

### تقسيم المشاكل إلى مكونات

نهجنا هو تقسيم المشاكل إلى أجزاء فردية (انظر أدناه) ثم تطبيق الذكاء الاصطناعي عليها واحدة تلو الأخرى. انظر أدناه لبعض الأمثلة.

<img width="2078" alt="augmented_challenges" src="https://github.com/danielmiessler/fabric/assets/50654/31997394-85a9-40c2-879b-b347e4701f06">

### الكثير من التروجات

التروجات جيدة لهذا، لكن أكبر تحد واجهته في عام 2023——والذي لا يزال موجودًا حتى اليوم—هو **العدد الهائل من تروجات الذكاء الاصطناعي الموجودة**. لدينا جميعًا تروجات مفيدة، لكن من الصعب اكتشاف تروجات جديدة، ومعرفة ما إذا كانت جيدة أم لا، _وإدارة إصدارات مختلفة من تلك التي نحبها._

إحدى الميزات الأساسية لـ `fabric` هي مساعدة الأشخاص في جمع وتكامل التروجات، والتي نسميها _الأنماط_، في أجزاء مختلفة من حياتهم.

يحتوي Fabric على أنماط لجميع أنواع أنشطة الحياة والعمل، بما في ذلك:

- استخراج الأجزاء الأكثر إثارة للاهتمام من مقاطع فيديو YouTube والبودكاست
- كتابة مقال بصوتك الخاص مع مجرد فكرة كمدخل
- تلخيص الأوراق الأكاديمية المعقدة
- إنشاء تروجات فنية للذكاء الاصطناعي متطابقة تمامًا لقطعة كتابة
- تقييم جودة المحتوى لمعرفة ما إذا كنت تريد قراءة/مشاهدة الشيء كله
- الحصول على ملخصات للمحتوى الطويل والممل
- شرح التعليمات البرمجية لك
- تحويل التوثيق السيئ إلى توثيق قابل للاستخدام
- إنشاء منشورات وسائل التواصل الاجتماعي من أي مدخل محتوى
- والمزيد بكثير…

## التثبيت

### تثبيت بسطر واحد (موصى به)

**Unix/Linux/macOS:**

```bash
curl -fsSL https://raw.githubusercontent.com/danielmiessler/fabric/main/scripts/installer/install.sh | bash
```

**Windows PowerShell:**

```powershell
iwr -useb https://raw.githubusercontent.com/danielmiessler/fabric/main/scripts/installer/install.ps1 | iex
```

> راجع [scripts/installer/README.md](./scripts/installer/README.md) لخيارات التثبيت المخصصة واستكشاف الأخطاء وإصلاحها.

### التنزيل اليدوي للثنائيات

يمكن العثور على أحدث أرشيفات الإصدار الثنائية وتوقيعات SHA256 المتوقعة لها في <https://github.com/danielmiessler/fabric/releases/latest>

### استخدام مديري الحزم

**ملاحظة:** استخدام Homebrew أو مديري حزم Arch Linux يجعل `fabric` متاحًا كـ `fabric-ai`، لذا أضف
الاسم المستعار التالي إلى ملفات بدء تشغيل shell الخاصة بك لمراعاة ذلك:

```bash
alias fabric='fabric-ai'
```

#### macOS (Homebrew)

`brew install fabric-ai`

#### Arch Linux (AUR)

`yay -S fabric-ai`

#### Windows

استخدم الأداة المدعومة رسميًا من Microsoft `Winget`:

`winget install danielmiessler.Fabric`

#### Windows (Scoop)

`scoop install fabric-ai`

### من المصدر

لتثبيت Fabric، [تأكد من تثبيت Go](https://go.dev/doc/install)، ثم قم بتشغيل الأمر التالي.

```bash
# تثبيت Fabric مباشرة من المستودع
go install github.com/danielmiessler/fabric/cmd/fabric@latest
```

### Docker

قم بتشغيل Fabric باستخدام صور Docker المبنية مسبقًا:

```bash
# استخدام أحدث صورة من Docker Hub
docker run --rm -it kayvan/fabric:latest --version

# استخدام إصدار محدد من GHCR
docker run --rm -it ghcr.io/ksylvan/fabric:v1.4.305 --version

# تشغيل الإعداد (لأول مرة)
mkdir -p $HOME/.fabric-config
docker run --rm -it -v $HOME/.fabric-config:/home/appuser/.config/fabric kayvan/fabric:latest --setup

# استخدام Fabric مع أنماطك
docker run --rm -it -v $HOME/.fabric-config:/home/appuser/.config/fabric kayvan/fabric:latest -p summarize

# تشغيل خادم واجهة برمجة التطبيقات REST (راجع قسم واجهة برمجة التطبيقات REST)
docker run --rm -it -p 8080:8080 -v $HOME/.fabric-config:/home/appuser/.config/fabric kayvan/fabric:latest --serve
```

**الصور المتاحة في:**

- Docker Hub: [kayvan/fabric](https://hub.docker.com/repository/docker/kayvan/fabric/general)
- GHCR: [ksylvan/fabric](https://github.com/ksylvan/fabric/pkgs/container/fabric)

راجع [scripts/docker/README.md](./scripts/docker/README.md) لبناء صور مخصصة والتكوين المتقدم.

### متغيرات البيئة

قد تحتاج إلى تعيين بعض متغيرات البيئة في `~/.bashrc` على Linux أو ملف `~/.zshrc` على mac لتتمكن من تشغيل أمر `fabric`. إليك مثال على ما يمكنك إضافته:

للحواسيب المبنية على Intel من Apple أو Linux

```bash
# متغيرات بيئة Golang
export GOROOT=/usr/local/go
export GOPATH=$HOME/go

# تحديث PATH لتشمل ثنائيات GOPATH و GOROOT
export PATH=$GOPATH/bin:$GOROOT/bin:$HOME/.local/bin:$PATH
```

للحواسيب المبنية على Apple Silicon من Apple

```bash
# متغيرات بيئة Golang
export GOROOT=$(brew --prefix go)/libexec
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$HOME/.local/bin:$PATH
```

### الإعداد

الآن قم بتشغيل الأمر التالي

```bash
# تشغيل الإعداد لإعداد أدلتك ومفاتيحك
fabric --setup
```

إذا كان كل شيء يعمل، فأنت جاهز للانطلاق.

### الترقية/الاستبدال

شيء رائع في Go هو أنه من السهل جدًا الترقية. ما عليك سوى تشغيل نفس الأمر الذي استخدمته للتثبيت في المقام الأول وستحصل دائمًا على أحدث إصدار.

```bash
go install github.com/danielmiessler/fabric/cmd/fabric@latest
```

إذا كنت تقوم باستبدال تثبيت موجود، فما عليك سوى تشغيل الأمر أعلاه وسيقوم Go تلقائيًا باستبدال الإصدار القديم بالإصدار الجديد.

**ملاحظة مهمة**: إذا قمت بتشغيل `fabric --setup` بعد الترقية، فقد تضطر إلى إعادة تكوين بعض الإعدادات أو المفاتيح إذا كان هناك تغييرات في التكوين بين الإصدارات.

### مزودو الذكاء الاصطناعي المدعومون

يدعم Fabric مجموعة واسعة من مزودي الذكاء الاصطناعي:

**التكاملات الأصلية:**

- OpenAI
- OpenAI Codex (اشتراك ChatGPT/Codex OAuth من خلال خلفية خاصة)
- Anthropic (Claude)
- Claude Code (اشتراك Claude من خلال `claude` CLI المحلي)
- Google Gemini
- Ollama (نماذج محلية)
- Azure OpenAI
- Amazon Bedrock
- Vertex AI
- LM Studio
- Perplexity

**المزودون المتوافقون مع OpenAI:**

- Abacus
- AIML
- API Route
- Apple Foundation Models (محلي، macOS 27 أو أحدث: قم بتشغيل `sudo fm license` مرة واحدة، ثم `fm serve --port 1976`؛ لا يوجد مفتاح API؛ حدده مرة واحدة في `fabric -S` لتمكينه)
- Cerebras
- Cheaper Inference
- DeepSeek
- DemonRoute
- DigitalOcean
- Eden AI
- GrokAI
- Groq
- Langdock
- LiteLLM
- MiniMax
- Mistral
- Novita AI
- OpenCode Go
- OpenCode Zen
- OpenRouter
- Opper
- OrcaRouter
- Pzero
- Requesty
- SiliconCloud
- Synthorai
- Together
- Venice AI
- Y-API
- Z AI

قم بتشغيل `fabric --setup` لتكوين مزودك المفضل (مزوديك)، أو استخدم `fabric --listvendors` لرؤية جميع البائعين المتاحين.

## الاستخدام

بمجرد إعداد كل شيء، إليك كيفية استخدامه.

```bash
fabric -h
```

لرؤية جميع الخيارات المتاحة والمساعدة.

## واجهة برمجة التطبيقات REST

يتضمن Fabric خادم واجهة برمجة تطبيقات REST مدمجًا يعرض جميع الوظائف الأساسية عبر HTTP. ابدأ الخادم بـ:

```bash
fabric --serve
```

يوفر الخادم نقاط نهاية لـ:

- إكمالات الدردشة مع ردود تدفقية
- إدارة الأنماط (إنشاء، قراءة، تحديث، حذف)
- إدارة السياق والجلسة
- قوائم النماذج والبائعين
- استخراج النصوص من YouTube
- إدارة التكوين

للحصول على وثائق نقطة النهاية الكاملة، وإعداد المصادقة، وأمثلة الاستخدام، راجع [وثائق REST API](docs/rest-api.md).

## أمثلة

> تستخدم الأمثلة التالية `pbpaste` الخاص بـ macOS للصق من الحافظة. راجع قسم [pbpaste](#pbpaste) أدناه للبدائل على Windows و Linux.

الآن دعونا نلقي نظرة على بعض الأشياء التي يمكنك القيام بها باستخدام Fabric.

1. تشغيل نمط `summarize` بناءً على الإدخال من `stdin`. في هذه الحالة، نص مقالة.

    ```bash
    pbpaste | fabric --pattern summarize
    ```

2. تشغيل نمط `analyze_claims` مع الخيار `--stream` للحصول على نتائج فورية وتدفقية.

    ```bash
    pbpaste | fabric --stream --pattern analyze_claims
    ```

3. تشغيل نمط `extract_wisdom` مع الخيار `--stream` للحصول على نتائج فورية وتدفقية من أي فيديو على Youtube (مشابه جدًا لمقطع الفيديو التقديمي الأصلي).

    ```bash
    fabric -y "https://youtube.com/watch?v=uXs-zPc63kM" --stream --pattern extract_wisdom
    ```

4. إنشاء أنماط - يجب إنشاء ملف .md مع النمط وحفظه في `~/.config/fabric/patterns/[yourpatternname]`.

5. تشغيل نمط `analyze_claims` على موقع ويب. يستخدم Fabric Jina AI لاستخراج محتوى الرابط إلى تنسيق markdown قبل إرساله إلى النموذج.

    ```bash
    fabric -u https://github.com/danielmiessler/fabric/ -p analyze_claims
    ```

## مجرد استخدام الأنماط

إذا كنت لا تبحث عن القيام بأي شيء متطور، وتريد فقط الكثير من التروجات الرائعة، يمكنك الانتقال إلى دليل [`/patterns`](https://github.com/danielmiessler/fabric/tree/main/data/patterns) والبدء في الاستكشاف!

نأمل أنه إذا لم تستخدم أي شيء آخر من Fabric، فإن الأنماط بحد ذاتها ستجعل المشروع مفيدًا.

يمكنك استخدام أي من الأنماط التي تراها هناك في أي تطبيق ذكاء اصطناعي لديك، سواء كان ChatGPT أو بعض التطبيقات أو المواقع الأخرى. خطتنا وتوقعنا هو أن الأشخاص سيقومون قريبًا بمشاركة المزيد بكثير من تلك التي نشرناها، وستكون أفضل بكثير من نماذجنا.

حكمة الجمهور للفوز.

## الأنماط المخصصة

قد ترغب في استخدام Fabric لإنشاء أنماطك المخصصة—لكن لا تشاركها مع الآخرين. لا مشكلة!

يدعم Fabric الآن دليل أنماط مخصص مخصص يحافظ على أنماطك الشخصية منفصلة عن الأنماط المدمجة. هذا يعني أن أنماطك المخصصة لن يتم الكتابة فوقها عند تحديث الأنماط المدمجة في Fabric.

### إعداد الأنماط المخصصة

1. قم بتشغيل إعداد Fabric:

   ```bash
   fabric --setup
   ```

2. حدد خيار "Custom Patterns" من قائمة الأدوات وأدخل مسار الدليل المطلوب (على سبيل المثال، `~/my-custom-patterns`)

3. سيقوم Fabric تلقائيًا بإنشاء الدليل إذا لم يكن موجودًا.

### استخدام الأنماط المخصصة

1. أنشئ هيكل دليل النمط المخصص الخاص بك:

   ```bash
   mkdir -p ~/my-custom-patterns/my-analyzer
   ```

2. أنشئ ملف النمط الخاص بك

   ```bash
   echo "أنت محلل خبير لـ ..." > ~/my-custom-patterns/my-analyzer/system.md
   ```

3. **استخدم النمط المخصص الخاص بك:**

   ```bash
   fabric --pattern my-analyzer "حلل هذا النص"
   ```

### كيف يعمل

- **نظام الأولوية**: تأخذ الأنماط المخصصة الأولوية على الأنماط المدمجة بنفس الاسم
- **التكامل السلس**: تظهر الأنماط المخصصة في `fabric --listpatterns` جنبًا إلى جنب مع الأنماط المدمجة
- **آمن للتحديث**: لا تتأثر أنماطك المخصصة أبدًا بـ `fabric --updatepatterns`
- **خاص افتراضيًا**: تظل أنماطك المخصصة خاصة تمامًا ما لم تشاركها صراحة

أنماطك المخصصة خاصة تمامًا ولن تتأثر بتحديثات Fabric!

## التطبيقات المساعدة

يجعل Fabric أيضًا استخدام بعض التطبيقات المساعدة الأساسية (الأدوات) لسهولة التكامل مع سير العمل المختلفة الخاصة بك. إليك بعض الأمثلة:

### `to_pdf`

`to_pdf` هو أمر مساعد يحول ملفات LaTeX إلى تنسيق PDF. يمكنك استخدامه هكذا:

```bash
to_pdf input.tex
```

سيؤدي هذا إلى إنشاء ملف PDF من ملف LaTeX المدخل في نفس الدليل.

يمكنك أيضًا استخدامه مع stdin الذي يعمل بشكل مثالي مع نمط `write_latex`:

```bash
echo "ai security primer" | fabric --pattern write_latex | to_pdf
```

سيؤدي هذا إلى إنشاء ملف PDF باسم `output.pdf` في الدليل الحالي.

### تثبيت `to_pdf`

لتثبيت `to_pdf`، قم بتثبيته بنفس طريقة تثبيت Fabric، فقط باسم مستودع مختلف.

```bash
go install github.com/danielmiessler/fabric/cmd/to_pdf@latest
```

تأكد من تثبيت توزيع LaTeX (مثل TeX Live أو MiKTeX) على نظامك، لأن `to_pdf` يتطلب `pdflatex` ليكون متاحًا في مسار نظامك.

### `code2context`

يستخدم `code2context` مع نمط `create_coding_feature`.
ينشئ تمثيل `json` لدليل التعليمات البرمجية الذي يمكن إدخاله إلى نموذج ذكاء اصطناعي
مع تعليمات لإنشاء ميزة جديدة أو تحرير التعليمات البرمجية بطريقة محددة.

راجع [README لنمط Create Coding Feature](./data/patterns/create_coding_feature/README.md) للتفاصيل.

قم بتثبيته أولاً باستخدام:

```bash
go install github.com/danielmiessler/fabric/cmd/code2context@latest
```

### `generate_changelog`

`generate_changelog` ينشئ سجلات التغيير من سجل commit في git وطلبات سحب GitHub. يمر عبر سجل git لمستودعك، ويستخرج معلومات PR، وينتج سجلات تغيير مكتوبة بشكل جيد.

```bash
generate_changelog --help
```

تشمل الميزات التخزين المؤقت SQLite للتحديثات التدريجية السريعة، وتكامل GraphQL API لـ GitHub لجلب PR فعال، وملخصات معززة بالذكاء الاصطناعي اختيارية باستخدام Fabric.

قم بتثبيته باستخدام:

```bash
go install github.com/danielmiessler/fabric/cmd/generate_changelog@latest
```

راجع [README لـ generate_changelog](./cmd/generate_changelog/README.md) لاستخدام مفصل والخيارات.

## pbpaste

[الأمثلة](#أمثلة) تستخدم برنامج macOS `pbpaste` للصق المحتوى من الحافظة إلى `fabric` كمدخل. `pbpaste` غير متوفر على Windows أو Linux، لكن هناك بدائل.

على Windows، يمكنك استخدام أمر PowerShell `Get-Clipboard` من موجه أمر PowerShell. إذا أردت، يمكنك أيضًا إعطاؤه اسمًا مستعارًا `pbpaste`. إذا كنت تستخدم PowerShell الكلاسيكي، قم بتحرير الملف `~\Documents\WindowsPowerShell\.profile.ps1`، أو إذا كنت تستخدم PowerShell Core، قم بتحرير `~\Documents\PowerShell\.profile.ps1` وأضف الاسم المستعار،

```powershell
Set-Alias pbpaste Get-Clipboard
```

على Linux، يمكنك استخدام `xclip -selection clipboard -o` للصق من الحافظة. من المحتمل أن تحتاج إلى تثبيت `xclip` باستخدام مدير الحزم الخاص بك. للأنظمة القائمة على Debian بما في ذلك Ubuntu،

```sh
sudo apt update
sudo apt install xclip -y
```

يمكنك أيضًا إنشاء اسم مستعار عن طريق تحرير `~/.bashrc` أو `~/.zshrc` وإضافة الاسم المستعار،

```sh
alias pbpaste='xclip -selection clipboard -o'
```

## واجهة الويب (تطبيق Fabric على الويب)

يتضمن Fabric الآن واجهة ويب مدمجة توفر بديل GUI لواجهة سطر الأوامر. راجع [README لتطبيق الويب](/web/README.md) لتعليمات التثبيت ونظرة عامة على الميزات.

## معلومات إضافية

> [!NOTE]
> شكر خاص للأشخاص التاليين لإلهامهم ومساهماتهم!

- _Jonathan Dunn_ لكونه مطور MVP المطلق في المشروع، بما في ذلك قيادة إصدار Go الجديد، وكذلك GUI! كل هذا أثناء كونه أيضًا طبيبًا بدوام كامل!
- _Caleb Sima_ لدفعي فوق حافة ما إذا كان يجب جعل هذا مشروعًا عامًا أم لا.
- _Eugen Eisler_ و _Frederick Ros_ لمساهماتهما القيمة في إصدار Go
- _David Peters_ لعمله على واجهة الويب.
- _Joel Parish_ لإدخاله المفيد للغاية على هيكل دليل Github للمشروع.
- _Joseph Thacker_ لفكرة العلم `-c` الذي يضيف سياقًا تم إنشاؤه مسبقًا في دليل `./config/fabric/` إلى جميع استعلامات النمط.
- _Jason Haddix_ لفكرة stitch (نمط متسلسل) لتصفية المحتوى باستخدام نموذج محلي قبل الإرسال إلى نموذج سحابي، أي تنظيف بيانات العملاء باستخدام `llama2` قبل الإرسال إلى `gpt-4` للتحليل.
- _Andre Guerra_ للمساعدة في العديد من المكونات لجعل الأشياء أبسط وأكثر قابلية للصيانة.

### المساهمون الأساسيون

<a href="https://github.com/danielmiessler"><img src="https://avatars.githubusercontent.com/u/50654?v=4" title="Daniel Miessler" width="50" height="50" alt="Daniel Miessler"></a>
<a href="https://github.com/xssdoctor"><img src="https://avatars.githubusercontent.com/u/9218431?v=4" title="Jonathan Dunn" width="50" height="50" alt="Jonathan Dunn"></a>
<a href="https://github.com/sbehrens"><img src="https://avatars.githubusercontent.com/u/688589?v=4" title="Scott Behrens" width="50" height="50" alt="Scott Behrens"></a>
<a href="https://github.com/agu3rra"><img src="https://avatars.githubusercontent.com/u/10410523?v=4" title="Andre Guerra" width="50" height="50" alt="Andre Guerra"></a>

### المساهمون

<a href="https://github.com/danielmiessler/fabric/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=danielmiessler/fabric" alt="contrib.rocks" />
</a>

مصنوع بـ [contrib.rocks](https://contrib.rocks).

`fabric` تم إنشاؤه بواسطة <a href="https://danielmiessler.com/subscribe" target="_blank">Daniel Miessler</a> في يناير 2024.
<br /><br />
<a href="https://twitter.com/intent/user?screen_name=danielmiessler">![X (formerly Twitter) Follow](https://img.shields.io/twitter/follow/danielmiessler)</a>

---

**💖 دعماً لهذا المشروع**

<div align="center">

<img src="https://img.shields.io/badge/Sponsor-❤️-EA4AAA?style=for-the-badge&logo=github-sponsors&logoColor=white" alt="Sponsor">

**أقضي مئات الساعات سنويًا على البرامج مفتوحة المصدر. إذا كنت ترغب في المساعدة في دعم هذا المشروع، يمكنك [رعايتي هنا](https://github.com/sponsors/danielmiessler). 🙏🏼**

</div>

---

## ترجمة
**HASAN ALDOY • Translate with Love from Bahrain • @aldoyh**

تمت ترجمة هذا المستند من الإنجليزية إلى العربية لمساعدة المجتمع العربي في استخدام إطار عمل Fabric الرائع. إذا وجدت أي خطأ في الترجمة أو لديك اقتراحات للتحسين، لا تتردد في التواصل.

العربية هي لغة غنية وجميلة، ودعمها في المشاريع التقنية يساعد في جعل التكنولوجيا أكثر شمولية وإتاحة للجميع.

🇧🇭 من البحرين مع الحب 💖