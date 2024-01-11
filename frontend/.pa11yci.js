require('dotenv').config();

const loginActions = [
  `set field [placeholder=''] to ${process.env.EMAIL}`,
  `set field [placeholder='Password'] to ${process.env.PASSWORD}`,
  "click element button[type='submit']",
];

module.exports = {
  urls: [
    {
      url: "http://localhost:3000",
    },
    {
      url: "http://localhost:3000/home",
      actions: [
        ...loginActions,
        "wait for path to be /home",
      ]
    },
    // {
    //   url: "http://localhost:3000/locations",
    //   actions: [
    //     ...loginActions,
    //     "navigate to http://localhost:3000/locations",
    //     "wait for path to be /locations",
    //     "wait for element html[lang='en'] to be added",
    //     "wait for element head > title to be added"
    //   ]
    // },
    // {
    //   url: "http://localhost:3000/profile",
    //   actions: [
    //     ...loginActions,
    //     "navigate to http://localhost:3000/profile",
    //     "wait for path to be /profile",
    //     "wait for element html[lang='en'] to be added",
    //     "wait for element head > title to be added"
    //   ]
    // },
    // {
    //   url: "http://localhost:3000/tools",
    //   actions: [
    //     ...loginActions,
    //     "navigate to http://localhost:3000/tools",
    //     "wait for path to be /tools",
    //     "wait for element html[lang='en'] to be added",
    //     "wait for element head > title to be added"
    //   ]
    // }
  ]
};