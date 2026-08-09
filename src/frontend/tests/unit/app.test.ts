import { describe, expect, it } from "vitest";
import { appName } from "../../src/index.js";

describe("media belajar anak bootstrap", () => {
  it("defines an application name", () => {
    expect(appName).toBe("Media Belajar Anak");
  });
});
