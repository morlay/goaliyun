package converter

import java.util.*
import java.util.regex.Pattern

/**
 * Handles singularization and pluralization.
 *
 * Largely taken from the defunct rogueweb project at
 * <link></link>http://code.google.com/p/rogueweb/
 * Original author Anthony Eden.
 */

class Inflection @JvmOverloads
constructor(private val pattern_: String, private val replacement_: String? = null, private val ignoreCase_: Boolean = true) {

    /**
     * Does the given word match?
     * @param word The word
     * @return True if it matches the inflection pattern
     */
    fun match(word: String): Boolean {
        val flags = if (ignoreCase_) Pattern.CASE_INSENSITIVE else 0
        return Pattern.compile(pattern_, flags).matcher(word).find()
    }

    /**
     * Replace the word with its pattern.
     * @param word The word
     * @return The result
     */
    fun replace(word: String): String {
        val flags = if (ignoreCase_) Pattern.CASE_INSENSITIVE else 0
        return Pattern.compile(pattern_, flags).matcher(word).replaceAll(replacement_)
    }

    companion object {
        private val plurals_ = ArrayList<Inflection>()
        private val singulars_ = ArrayList<Inflection>()
        private val uncountables_ = ArrayList<String>()

        init {
            // plural is "singular to plural form"
            // singular is "plural to singular form"
            plural("$", "s")
            plural("s$", "s")
            plural("(ax|test)is$", "$1es")
            plural("(octop|vir)us$", "$1i")
            plural("(alias|status)$", "$1es")
            plural("(bu)s$", "$1ses")
            plural("(buffal|tomat)o$", "$1oes")
            plural("([ti])um$", "$1a")
            plural("sis$", "ses")
            plural("(?:([^f])fe|([lr])f)$", "$1$2ves")
            plural("(hive)$", "$1s")
            plural("([^aeiouy]|qu)y$", "$1ies")
            // plural("([^aeiouy]|qu)ies$", "$1y");
            plural("(x|ch|ss|sh)$", "$1es")
            plural("(matr|vert|ind)ix|ex$", "$1ices")
            plural("([m|l])ouse$", "$1ice")
            plural("^(ox)$", "$1en")
            plural("(quiz)$", "$1zes")

            singular("s$", "")
            singular("(n)ews$", "$1ews")
            singular("([ti])a$", "$1um")
            singular("((a)naly|(b)a|(d)iagno|(p)arenthe|(p)rogno|(s)ynop|(t)he)ses$", "$1$2sis")
            singular("(^analy)ses$", "$1sis")
            singular("([^f])ves$", "$1fe")
            singular("(hive)s$", "$1")
            singular("(tive)s$", "$1")
            singular("([lr])ves$", "$1f")
            singular("([^aeiouy]|qu)ies$", "$1y")
            singular("(s)eries$", "$1eries")
            singular("(m)ovies$", "$1ovie")
            singular("(x|ch|ss|sh)es$", "$1")
            singular("([m|l])ice$", "$1ouse")
            singular("(bus)es$", "$1")
            singular("(o)es$", "$1")
            singular("(shoe)s$", "$1")
            singular("(cris|ax|test)es$", "$1is")
            singular("(octop|vir)i$", "$1us")
            singular("(alias|status)es$", "$1")
            singular("^(ox)en", "$1")
            singular("(vert|ind)ices$", "$1ex")
            singular("(matr)ices$", "$1ix")
            singular("(quiz)zes$", "$1")

            irregular("person", "people")
            irregular("man", "men")
            irregular("child", "children")
            irregular("sex", "sexes")
            irregular("move", "moves")

            uncountable("equipment")
            uncountable("information")
            uncountable("rice")
            uncountable("money")
            uncountable("species")
            uncountable("series")
            uncountable("fish")
            uncountable("sheep")
        }

        private fun plural(pattern: String, replacement: String) {
            plurals_.add(0, Inflection(pattern, replacement))
        }

        private fun singular(pattern: String, replacement: String) {
            singulars_.add(0, Inflection(pattern, replacement))
        }

        private fun irregular(s: String, p: String) {
            plural("(" + s.substring(0, 1) + ")" + s.substring(1) + "$", "$1" + p.substring(1))
            singular("(" + p.substring(0, 1) + ")" + p.substring(1) + "$", "$1" + s.substring(1))
        }

        private fun uncountable(word: String) {
            uncountables_.add(word)
        }

        /**
         * Return the pluralized version of a word.
         * @param word The word
         * @return The pluralized word
         */
        fun pluralize(word: String): String {
            if (isUncountable(word)) {
                return word
            } else {
                for (inflection in plurals_) {
                    if (inflection.match(word)) {
                        return inflection.replace(word)
                    }
                }
                return word
            }
        }

        /**
         * Return the singularized version of a word.
         * @param word The word
         * @return The singularized word
         */
        fun singularize(word: String): String {
            if (isUncountable(word)) {
                return word
            } else {
                for (inflection in singulars_) {
                    if (inflection.match(word)) {
                        return inflection.replace(word)
                    }
                }
            }
            return word
        }

        /**
         * Return true if the word is uncountable.
         * @param word The word
         * @return True if it is uncountable
         */
        fun isUncountable(word: String): Boolean {
            for (w in uncountables_) {
                if (w.equals(word, ignoreCase = true)) {
                    return true
                }
            }
            return false
        }
    }

}