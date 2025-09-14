
package main;

public interface ConfigLoader<T> {
    T load(String filePath);
}
