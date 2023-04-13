import type { MetaFunction } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import { useState } from "react";
import Footer from "~/components/landingPage/Footer";
import Navbar from "~/components/landingPage/Navbar";
import { fetchGenerativeAITools } from "./api/auth.server";
import Tabs from "react-simply-tabs";

export const meta: MetaFunction = () => ({
  charset: "utf-8",
  title: "Generative AI Tools | PureML",
  viewport: "width=device-width,initial-scale=1",
});

export async function loader() {
  const toolsList = await fetchGenerativeAITools();
  return toolsList;
}

export default function MLTools() {
  const data = useLoaderData();
  const [activeTabIndex, setActiveTabIndex] = useState(0);
  return (
    <div className="bg-slate-50 landingpg-font flex flex-col justify-center">
      <div className="generativeAiToolsbg bg-cover">
        <Navbar />
        <div className="flex flex-col justify-center md:items-center md:text-center gap-y-4 px-4 md:px-0 h-80 md:h-96 lg:h-[32rem]">
          <h1 className="font-bold text-4xl md:text-5xl lg:text-6xl">
            Generative AI Tools
          </h1>
          <h1 className="text-slate-600 text-lg md:text-2xl lg:text-3xl md:w-4/5 lg:w-3/5 xl:w-2/5 2xl:w-1/5">
            A web page listing various AI tools that can generate new and unique
            content, such as text, images, and music.
          </h1>
        </div>
      </div>
      <div className="border-b border-slate-200 md:h-4/5">
        <div className="bg-slate-50 flex justify-center">
          <div className="w-full md:max-w-screen-xl px-4 md:px-8">
            <Tabs
              activeTabIndex={activeTabIndex}
              onRequestChange={setActiveTabIndex}
              wrapperProps={{ style: { display: "flex" } }}
              tabsWrapperProps={{ style: { width: "100%" } }}
              controls={[
                <div
                  key="title"
                  className="font-medium text-base md:text-2xl bg-gradient-to-r from-sky-100 p-4 md:py-2.5 md:px-6"
                >
                  GENERATIONS
                </div>,
                <button
                  type="button"
                  className="text-base md:text-2xl text-slate-400 mt-6 px-2 md:py-2 md:px-6"
                  key="tab1"
                >
                  TEXT
                </button>,
                <button
                  type="button"
                  className="text-base md:text-2xl text-slate-400 py-2 px-2 md:px-6"
                  key="tab2"
                >
                  IMAGE
                </button>,
                <button
                  type="button"
                  className="text-base md:text-2xl text-slate-400 py-2 px-2 md:px-6"
                  key="tab3"
                >
                  MUSIC
                </button>,
                <button
                  type="button"
                  className="text-base md:text-2xl text-slate-400 py-2 px-2 md:px-6"
                  key="tab4"
                >
                  VIDEO
                </button>,
                <button
                  type="button"
                  className="text-base md:text-2xl text-slate-400 py-2 px-2 md:px-6"
                  key="tab5"
                >
                  SPEECH
                </button>,
                <button
                  type="button"
                  className="text-base md:text-2xl text-slate-400 py-2 px-2 md:px-6"
                  key="tab6"
                >
                  CHATBOX
                </button>,
                <button
                  type="button"
                  className="text-base md:text-2xl text-slate-400 py-2 px-2 md:px-6"
                  key="tab7"
                >
                  GAME
                </button>,
                <button
                  type="button"
                  className="text-base md:text-2xl text-slate-400 py-2 px-2 md:px-6"
                  key="tab8"
                >
                  CODE
                </button>,
              ]}
              controlsWrapperProps={{
                style: {
                  width: "30%",
                  paddingBottom: "5vh",
                  paddingTop: "5vh",
                  borderRight: "1px solid #e2e8f0",
                  display: "flex",
                  flexDirection: "column",
                  alignItems: "start",
                },
              }}
              activeControlProps={{
                className: "active",
                style: { color: "#020617" },
              }}
            >
              {data ? (
                <div className="px-6 py-[5vh] md:px-12 h-full">
                  <div className="grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-12 min-w-72">
                    {data.map((tool: any, key: number) => (
                      <>
                        {tool.properties.Category.multi_select.some(
                          (item: any) => item.name === "Text"
                        ) && (
                          <a
                            href={tool.properties.Link.url}
                            key={key}
                            className="flex flex-col gap-y-4 p-4 text-slate-600 rounded-lg hover:bg-slate-100"
                          >
                            <div className="flex gap-x-4 items-center">
                              <div className="bg-slate-100 w-10 h-10 flex justify-center items-center rounded-lg border border-slate-200">
                                <img
                                  src={
                                    tool.properties.Logo.files[0].external.url
                                  }
                                  alt="Logo"
                                  className="w-7 h-7"
                                />
                              </div>
                              <div className="text-2xl font-medium">
                                {tool.properties.Name.title[0].plain_text}
                              </div>
                            </div>
                            <div className="flex flex-col gap-y-2">
                              <span className="text-lg">
                                {
                                  tool.properties.Description.rich_text[0]
                                    .plain_text
                                }
                              </span>
                            </div>
                          </a>
                        )}
                      </>
                    ))}
                  </div>
                </div>
              ) : (
                "Tools coming soon..."
              )}
              {data ? (
                <div className="px-6 py-8 md:px-12 md:py-16 h-full">
                  <div className="grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-12 min-w-72">
                    {data.map((tool: any, key: number) => (
                      <>
                        {tool.properties.Category.multi_select.some(
                          (item: any) => item.name === "Text"
                        ) && (
                          <a
                            href={tool.properties.Link.url}
                            key={key}
                            className="flex flex-col gap-y-4 p-4 text-slate-600 rounded-lg hover:bg-slate-100"
                          >
                            <div className="flex gap-x-4 items-center">
                              <div className="bg-slate-100 w-10 h-10 flex justify-center items-center rounded-lg border border-slate-200">
                                <img
                                  src={
                                    tool.properties.Logo.files[0].external.url
                                  }
                                  alt="Logo"
                                  className="w-7 h-7"
                                />
                              </div>
                              <div className="text-2xl font-medium">
                                {tool.properties.Name.title[0].plain_text}
                              </div>
                            </div>
                            <div className="flex flex-col gap-y-2">
                              <span className="text-lg">
                                {
                                  tool.properties.Description.rich_text[0]
                                    .plain_text
                                }
                              </span>
                            </div>
                          </a>
                        )}
                      </>
                    ))}
                  </div>
                </div>
              ) : (
                "Tools coming soon..."
              )}
              {data ? (
                <div className="px-6 py-8 md:px-12 md:py-16 h-full">
                  <div className="grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-12 min-w-72">
                    {data.map((tool: any, key: number) => (
                      <>
                        {tool.properties.Category.multi_select.some(
                          (item: any) => item.name === "Image"
                        ) && (
                          <a
                            href={tool.properties.Link.url}
                            key={key}
                            className="flex flex-col gap-y-4 p-4 text-slate-600 rounded-lg hover:bg-slate-100"
                          >
                            <div className="flex gap-x-4 items-center">
                              <div className="bg-slate-100 w-10 h-10 flex justify-center items-center rounded-lg border border-slate-200">
                                <img
                                  src={
                                    tool.properties.Logo.files[0].external.url
                                  }
                                  alt="Logo"
                                  className="w-7 h-7"
                                />
                              </div>
                              <div className="text-2xl font-medium">
                                {tool.properties.Name.title[0].plain_text}
                              </div>
                            </div>
                            <div className="flex flex-col gap-y-2">
                              <span className="text-lg">
                                {
                                  tool.properties.Description.rich_text[0]
                                    .plain_text
                                }
                              </span>
                            </div>
                          </a>
                        )}
                      </>
                    ))}
                  </div>
                </div>
              ) : (
                "Tools coming soon..."
              )}
              {data ? (
                <div className="px-6 py-8 md:px-12 md:py-16 h-full">
                  <div className="grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-12 min-w-72">
                    {data.map((tool: any, key: number) => (
                      <>
                        {tool.properties.Category.multi_select.some(
                          (item: any) => item.name === "Music"
                        ) && (
                          <a
                            href={tool.properties.Link.url}
                            key={key}
                            className="flex flex-col gap-y-4 p-4 text-slate-600 rounded-lg hover:bg-slate-100"
                          >
                            <div className="flex gap-x-4 items-center">
                              <div className="bg-slate-100 w-10 h-10 flex justify-center items-center rounded-lg border border-slate-200">
                                <img
                                  src={
                                    tool.properties.Logo.files[0].external.url
                                  }
                                  alt="Logo"
                                  className="w-7 h-7"
                                />
                              </div>
                              <div className="text-2xl font-medium">
                                {tool.properties.Name.title[0].plain_text}
                              </div>
                            </div>
                            <div className="flex flex-col gap-y-2">
                              <span className="text-lg">
                                {
                                  tool.properties.Description.rich_text[0]
                                    .plain_text
                                }
                              </span>
                            </div>
                          </a>
                        )}
                      </>
                    ))}
                  </div>
                </div>
              ) : (
                "Tools coming soon..."
              )}
              {data ? (
                <div className="px-6 py-8 md:px-12 md:py-16 h-full">
                  <div className="grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-12 min-w-72">
                    {data.map((tool: any, key: number) => (
                      <>
                        {tool.properties.Category.multi_select.some(
                          (item: any) => item.name === "Video"
                        ) && (
                          <a
                            href={tool.properties.Link.url}
                            key={key}
                            className="flex flex-col gap-y-4 p-4 text-slate-600 rounded-lg hover:bg-slate-100"
                          >
                            <div className="flex gap-x-4 items-center">
                              <div className="bg-slate-100 w-10 h-10 flex justify-center items-center rounded-lg border border-slate-200">
                                <img
                                  src={
                                    tool.properties.Logo.files[0].external.url
                                  }
                                  alt="Logo"
                                  className="w-7 h-7"
                                />
                              </div>
                              <div className="text-2xl font-medium">
                                {tool.properties.Name.title[0].plain_text}
                              </div>
                            </div>
                            <div className="flex flex-col gap-y-2">
                              <span className="text-lg">
                                {
                                  tool.properties.Description.rich_text[0]
                                    .plain_text
                                }
                              </span>
                            </div>
                          </a>
                        )}
                      </>
                    ))}
                  </div>
                </div>
              ) : (
                "Tools coming soon..."
              )}
              {data ? (
                <div className="px-6 py-8 md:px-12 md:py-16 h-full">
                  <div className="grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-12 min-w-72">
                    {data.map((tool: any, key: number) => (
                      <>
                        {tool.properties.Category.multi_select.some(
                          (item: any) => item.name === "Speech"
                        ) && (
                          <a
                            href={tool.properties.Link.url}
                            key={key}
                            className="flex flex-col gap-y-4 p-4 text-slate-600 rounded-lg hover:bg-slate-100"
                          >
                            <div className="flex gap-x-4 items-center">
                              <div className="bg-slate-100 w-10 h-10 flex justify-center items-center rounded-lg border border-slate-200">
                                <img
                                  src={
                                    tool.properties.Logo.files[0].external.url
                                  }
                                  alt="Logo"
                                  className="w-7 h-7"
                                />
                              </div>
                              <div className="text-2xl font-medium">
                                {tool.properties.Name.title[0].plain_text}
                              </div>
                            </div>
                            <div className="flex flex-col gap-y-2">
                              <span className="text-lg">
                                {
                                  tool.properties.Description.rich_text[0]
                                    .plain_text
                                }
                              </span>
                            </div>
                          </a>
                        )}
                      </>
                    ))}
                  </div>
                </div>
              ) : (
                "Tools coming soon..."
              )}
              {data ? (
                <div className="px-6 py-8 md:px-12 md:py-16 h-full">
                  <div className="grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-12 min-w-72">
                    {data.map((tool: any, key: number) => (
                      <>
                        {tool.properties.Category.multi_select.some(
                          (item: any) => item.name === "Chatbot"
                        ) && (
                          <a
                            href={tool.properties.Link.url}
                            key={key}
                            className="flex flex-col gap-y-4 p-4 text-slate-600 rounded-lg hover:bg-slate-100"
                          >
                            <div className="flex gap-x-4 items-center">
                              <div className="bg-slate-100 w-10 h-10 flex justify-center items-center rounded-lg border border-slate-200">
                                <img
                                  src={
                                    tool.properties.Logo.files[0].external.url
                                  }
                                  alt="Logo"
                                  className="w-7 h-7"
                                />
                              </div>
                              <div className="text-2xl font-medium">
                                {tool.properties.Name.title[0].plain_text}
                              </div>
                            </div>
                            <div className="flex flex-col gap-y-2">
                              <span className="text-lg">
                                {
                                  tool.properties.Description.rich_text[0]
                                    .plain_text
                                }
                              </span>
                            </div>
                          </a>
                        )}
                      </>
                    ))}
                  </div>
                </div>
              ) : (
                "Tools coming soon..."
              )}
              {data ? (
                <div className="px-6 py-8 md:px-12 md:py-16 h-full">
                  <div className="grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-12 min-w-72">
                    {data.map((tool: any, key: number) => (
                      <>
                        {tool.properties.Category.multi_select.some(
                          (item: any) => item.name === "Game"
                        ) && (
                          <a
                            href={tool.properties.Link.url}
                            key={key}
                            className="flex flex-col gap-y-4 p-4 text-slate-600 rounded-lg hover:bg-slate-100"
                          >
                            <div className="flex gap-x-4 items-center">
                              <div className="bg-slate-100 w-10 h-10 flex justify-center items-center rounded-lg border border-slate-200">
                                <img
                                  src={
                                    tool.properties.Logo.files[0].external.url
                                  }
                                  alt="Logo"
                                  className="w-7 h-7"
                                />
                              </div>
                              <div className="text-2xl font-medium">
                                {tool.properties.Name.title[0].plain_text}
                              </div>
                            </div>
                            <div className="flex flex-col gap-y-2">
                              <span className="text-lg">
                                {
                                  tool.properties.Description.rich_text[0]
                                    .plain_text
                                }
                              </span>
                            </div>
                          </a>
                        )}
                      </>
                    ))}
                  </div>
                </div>
              ) : (
                "Tools coming soon..."
              )}
              {data ? (
                <div className="px-6 py-8 md:px-12 md:py-16 h-full">
                  <div className="grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-12 min-w-72">
                    {data.map((tool: any, key: number) => (
                      <>
                        {tool.properties.Category.multi_select.some(
                          (item: any) => item.name === "Code"
                        ) && (
                          <a
                            href={tool.properties.Link.url}
                            key={key}
                            className="flex flex-col gap-y-4 p-4 text-slate-600 rounded-lg hover:bg-slate-100"
                          >
                            <div className="flex gap-x-4 items-center">
                              <div className="bg-slate-100 w-10 h-10 flex justify-center items-center rounded-lg border border-slate-200">
                                <img
                                  src={
                                    tool.properties.Logo.files[0].external.url
                                  }
                                  alt="Logo"
                                  className="w-7 h-7"
                                />
                              </div>
                              <div className="text-2xl font-medium">
                                {tool.properties.Name.title[0].plain_text}
                              </div>
                            </div>
                            <div className="flex flex-col gap-y-2">
                              <span className="text-lg">
                                {
                                  tool.properties.Description.rich_text[0]
                                    .plain_text
                                }
                              </span>
                            </div>
                          </a>
                        )}
                      </>
                    ))}
                  </div>
                </div>
              ) : (
                "Tools coming soon..."
              )}
            </Tabs>
          </div>
        </div>
      </div>
      <Footer />
    </div>
  );
}
