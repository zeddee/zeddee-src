---
title: "Zed Tan"
---
<style>
.this-is-zed {
display: inline-block;
vertical-align: top;
width:10em;
height:10em;
border-radius:50%;
-webkit-border-radius:50%;
-moz-border-radius:50%;
overflow:hidden;
  transition-property:box-shadow;
  transition-duration:300ms;
  box-shadow: 0.3em 0.3em 1em hsla(0,100%,0%,.2);
  -moz-box-shadow: 0.3em 0.3em 1em hsla(0,100%,0%,.2);
  -webkit-box-shadow: 0.3em 0.3em 1em hsla(0,100%,0%,.2);
}
.this-is-zed:hover {
  box-shadow: 0.3em 0.3em 1em hsla(0,100%,0%,.4);
  -moz-box-shadow: 0.3em 0.3em 1em hsla(0,100%,0%,.4);
  -webkit-box-shadow: 0.3em 0.3em 1em hsla(0,100%,0%,.4);
}
.this-is-zed>img{
  width:100%;
}
.cvDescription{
  vertical-align: top;
  display: inline-block;
  width: 60%;
  margin-right: 2em;
}
@media only screen and (max-width: 532px) {
  .cvDescription {
    width: 100%;
  }
}
</style>
<div class="cvHeader">
<div class="cvDescription">
Zed is a former educator, and documentarian by nature (i.e. asks a lot of questions and writes stuff down).

He's happiest when building content systems and the teams that manage them.

You can reach him at [zed@shootbird.work](mailto:zed@shootbird.work)
</div>
<div class="this-is-zed">
<a href="mailto:zed@shootbird.work" alt="zed@shootbird.work"><img src="/images/zed.jpg" /></a>
</div>
</div>

## Skills

- **Teaching/Lecturing/Presentations**
  - Taught design theory in vocational institutes for 3 years
  - Taught various introductory courses at Nanyang Technological University for 2 years
- **Writing/Editing**
  - Technical writer since 2015.
  - Build and maintain SSG-based toolchains.
- **HTML/CSS/JS**
  - [flare-tocgen](https://github.com/zeddee/flare-tocgen)
    (archived), still in use at
    [docs.groundlabs.com](https://www.groundlabs.com/documentation/er/Content/Release-Notes.html)
    in 2021.
  - [vuepress-and-xlsx](https://github.com/zeddee/vuepress-and-xlsx)
    (archived) where I use a NodeJS stack
    to generate a Vuepress site from excel spreadsheets.
- **Go/Golang**
  - Wrote Go guides for MessageBird (see **Published Work**)
  - Minor contributions to MessageBird go-rest-api:
    - [implemented onFinish attribute for CallFlowRecordStep #54](https://github.com/messagebird/go-rest-api/pull/54)
    - [Export links field #56](https://github.com/messagebird/go-rest-api/pull/56)
  - Small tools:
    - [clear-archiver](https://github.com/zeddee/clear-archiver)
- **Python**
  - [bad-json-to-rst-tables](https://github.com/zeddee/bad-json-to-rst-tables)
  - Minor contribution to [sphinx-contrib/confluencebuilder](https://github.com/sphinx-contrib/confluencebuilder/pull/414)
  - [python-ast-docgen](https://github.com/zeddee/python-ast-docgen)
  - [parsedocinfo](https://github.com/zeddee/parsedocinfo)
- **\*NIX**
  - Various [techdoc_utility_scripts](https://github.com/zeddee/techdocs_utility_scripts)
  - Minor contribution to [things.sh](https://github.com/AlexanderWillner/things.sh/pull/1)
- **CI**
  - Build and maintain GitLab pipelines that produce https://docs.eclecticiq.com
  - Ansible scripts



## Published Work

**Technical Content for the [Temasys.io Blog](https://temasys.io)**

- ["Ripping off the band-aid: Chrome’s shift to 'Unified plan'"](https://temasys.io/ripping-off-the-band-aid-chromes-shift-to-unified-plan/)

**Blockchain with Go guides for Consensys/Kauri** ([github repo](https://github.com/kauri-io/Content/tree/master/Go/write-basic-quiz-dapp-in-go/quiz-dapp)):

- [Creating a DApp in Go with Geth](https://kauri.io/#collections/A%20Hackathon%20Survival%20Guide/creating-a-dapp-in-go-with-geth/)


[**Ground Labs Documentation**](https://docs.groundlabs.com/)

- See [work](#work).


**MessageBird Developer Guides** (samples):

- [Number Masking Guide (Go/Golang)](https://github.com/messagebirdguides/masked-numbers-guide-go)
- [SMS Marketing Subscriptions (Go/Golang)](https://github.com/messagebirdguides/subscriptions-guide-go)

## Work

### **Senior Technical Writer** at [EclecticIQ](https://docs.eclecticiq.com)

* Documentation site: https://docs.eclecticiq.com
* Embedded in integrations team.
* Work closely with frontend and backend teams.
* Site migration from
  [Confluence](https://www.atlassian.com/software/confluence)
  to [Sphinx](https://www.sphinx-doc.org)-based site.

### **Technical Writer** at [AI Singapore](https://www.aisingapore.org)

* Worked on web site information architecture for [Makerspace](https://makerspace.aisingapore.org)
* Worked on technical documentation for AI Speech Lab.
* Worked on internal knowledge and information sharing system for AI Singapore's engineering team.
Built a semi-automated pipeline that would publish project documentation hosted on GitLab repositories
to a URL on the intranet.

### **Contract Technical Writer** at [Temasys Communications](https://temasys.io)

* Wrote technical content for blogs and marketing material.
* Worked closely with development team to communicate releases, technical developments, and product updates.
* Helped make case and set the groundwork for dedicated technical documentation efforts.
* Worked closely with development team to document a pre-alpha product, which involved digging into C++ source.

### [**Ground Labs Documentation**](https://docs.groundlabs.com/)

* Lead technical writer. Responsible for keeping documentation site up to date and identifying informational gaps.
* Established beach-head for documentation.
* Embedded in engineering team.
* Site migration from [Confluence](https://www.atlassian.com/software/confluence) to [MadCap Flare](https://www.madcapsoftware.com/products/flare/).

Documentation sets:

- Enterprise Recon Documentation[(pdf)](/pdfs/ER-2.0.26-documentation.pdf)
- Card Recon Documentation[(pdf)](/pdfs/CR-2.0.25-documentation.pdf)
- Data Recon Documentation[(pdf)](/pdfs/DR-2.0.25-documentation.pdf)

EU GDPR-related Research Project

* Project management.
* Managed 20 freelancers gathering research from 28 EU countries.
* Cataloged and documented research for developer use.

## Experience

**Technical Writer (Tech lead)** at [EclecticIQ](https://ecelcticiq.com) (Feb 2020 – present)

**Technical Writer** at [AI Singapore](https://aisingapore.org) (October 2019 – Feb 2020)

**Technical Content Writer (Contract)** at [Temasys Communications](https://temasys.io) (Feb – May 2019)

**Technical Writer** at [Ground Labs](https://groundlabs.com) (Feb 2016 – Jul 2018)

**Adjunct Instructor** at

* Singapore Polytechnic, _Theory of Design_ (Apr 2014 – Feb 2016)
* NTU WKWSCI[^1], _Introduction to Film Studies_ (Aug 2013 – Dec 2013)
* NTU WKWSCI, _Web Design and Technologies_ (Jan 2015 – Dec 2015)

[^1]: Nanyang Technological University Wee Kim Wee School of Communication and Information

## Other

"What _Time_ Got Wrong About _The Last of Us_", _Kill Screen_, published 8 Oct 2014, available: [https://killscreen.com/articles/what-time-got-wrong-about-last-us/](https://killscreen.com/articles/what-time-got-wrong-about-last-us/)

Guest Lecture, The Documentary Genre and Game Narratives ([presentation workbook](/pdfs/CS4026-Documentaries-Games-and-Narratives.pdf)), 18 Oct 2016.
