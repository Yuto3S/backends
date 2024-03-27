package com.example.demo

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.data.annotation.Id
import org.springframework.data.relational.core.mapping.Table
import org.springframework.data.repository.CrudRepository
import org.springframework.jdbc.core.JdbcTemplate
import org.springframework.stereotype.Service
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.PostMapping
import org.springframework.web.bind.annotation.RequestBody
import org.springframework.web.bind.annotation.RequestParam
import org.springframework.web.bind.annotation.RestController
import java.util.*
import org.springframework.jdbc.core.query
import org.springframework.web.bind.annotation.PathVariable
import kotlin.jvm.optionals.toList

@SpringBootApplication
class DemoApplication

fun main(args: Array<String>) {
	runApplication<DemoApplication>(*args)
}

@RestController
class MessageController(val service: MessageService) {
	@GetMapping("/")
	fun index(@RequestParam("name") name: String) = "Hello, $name!"

	@GetMapping("/listof")
	fun index2() = listOf(
		Message("1", "Hello"),
		Message("2", "Bonjour"),
		Message("3", "Privet!"),
	)

	@GetMapping("/messages")
	fun findMessages(): List<Message> = service.findMessages()

	@GetMapping("messages/{id}")
	fun getMessageById(@PathVariable id: String): List<Message> = service.findMessageById(id)

	@PostMapping("/messages")
	fun postMessages(@RequestBody message:Message) {
		service.save(message)
	}
}

@Table("MESSAGES")
data class Message(@Id val id: String?, val text: String)

interface MessageRepository : CrudRepository<Message, String>

@Service
class MessageService(val db: MessageRepository) {
	fun findMessages(): List<Message> = db.findAll().toList()

	fun findMessageById(id: String): List<Message> = db.findById(id).toList()

	fun save(message: Message) {
		db.save(message)
	}

	fun <T : Any> Optional<out T>.toList(): List<T> = if (isPresent) listOf(get()) else emptyList()
}